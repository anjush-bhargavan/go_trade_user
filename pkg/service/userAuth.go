package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/anjush-bhargavan/go_trade_user/config"
	"github.com/anjush-bhargavan/go_trade_user/pkg/model"
	pb "github.com/anjush-bhargavan/go_trade_user/pkg/proto"
	"github.com/anjush-bhargavan/go_trade_user/utils"
	"gorm.io/gorm"
)

// SignupService method receive the data in proto messages and start the verfification process.
func (u *UserService) SignupService(p *pb.Signup) (*pb.Response, error) {
	hashedPass, err := utils.HashPassword(p.Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in hashing password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("unable to hash password")
	}

	user := &model.User{
		UserName:     p.User_Name,
		Email:        p.Email,
		Password:     hashedPass,
		Mobile:       p.Mobile,
		ReferralCode: p.Referral_Code,
	}

	_, err = u.Repo.FindUserByEmail(user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "email already exists",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	resp, err := u.twilio.SendTwilioOTP(p.Mobile)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in sending otp using twilio",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	if resp.Status != nil {
		fmt.Println(*resp.Status)
	} else {
		fmt.Println(resp.Status)
	}

	userData, err := json.Marshal(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in marshaling data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error while marshaling data")
	}

	key := fmt.Sprintf("user_%v", p.Email)
	err = u.redis.SetDataInRedis(key, userData, time.Minute*3)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in setting data in redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error setting data in redis")
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Go to Verification page",
	}, nil
}

// VerificationService handles the verification process and proceed with signup
func (u *UserService) VerificationService(p *pb.OTP) (*pb.Response, error) {

	key := fmt.Sprintf("user_%v", p.Email)

	userData, err := u.redis.GetFromRedis(key)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting data from redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	//Unmarshal the data from redis
	var user model.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in unmarshaling data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	err = u.twilio.VerifyTwilioOTP(user.Mobile, p.OTP)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in verfying twilio otp",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	referralCode := user.ReferralCode

	newReferralCode := ""
	for newReferralCode == "" {
		code := utils.GenerateReferralCode(6)
		_, err := u.Repo.FindReferralCode(code)
		if err == gorm.ErrRecordNotFound {
			newReferralCode = code
		} else if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "error in creating referralcode",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, err
		}
	}
	user.ReferralCode = newReferralCode

	userID, err := u.Repo.CreateUser(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in creating user in database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("unable to create user")
	}
	if referralCode != "" {
		referrer, err := u.Repo.FindReferralCode(referralCode)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "Invalid referral code",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, err
		}
		subject1 := fmt.Sprintf("Referral bonus for user: %d", userID)
		transactionReferrer := model.Transaction{
			UserID: referrer.ID,
			Name:   subject1,
			Amount: 100,
		}
		err = u.Repo.CreateTransaction(&transactionReferrer)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "Error creating transaction",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, err
		}
		subject2 := fmt.Sprintf("Referral bonus by user: %d", referrer.ID)
		transactionReferral := model.Transaction{
			UserID: userID,
			Name:   subject2,
			Amount: 100,
		}
		err = u.Repo.CreateTransaction(&transactionReferral)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "Error creating transaction",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, err
		}
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "User created successfully",
	}, nil
}

// LoginService handles the logging in of user it check the password and generate token
func (u *UserService) LoginService(p *pb.Login) (*pb.Response, error) {
	user, err := u.Repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(p.Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password is incorrect",
		}, errors.New("password incorrect")
	}

	if user.IsBlocked {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user is blocked by admin",
		}, errors.New("you are blocked by the admin")
	}

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECRETKEY, user.Email, user.ID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in generating token",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error generating token")
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Login successful",
		Payload: &pb.Response_Data{Data: jwtToken},
	}, nil
}
