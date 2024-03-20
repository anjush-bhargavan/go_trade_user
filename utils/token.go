package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

//Claims struct defines the variables that jwt token holds
type Claims struct {
	UserID uint
	Email  string
	jwt.StandardClaims
}

//GenerateToken will generate token for 5 hours with given data
func GenerateToken(key, email string, userID uint) (string, error) {
	expTime := time.Now().Add(time.Hour * 5).Unix()

	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v",email,err.Error())
		return "", err
	}
	return signedToken, nil
}
