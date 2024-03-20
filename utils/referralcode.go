package utils

import (
	"math/rand"
	"time"
)


// GenerateReferralCode will take length and returns a random string
func GenerateReferralCode(length int) string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.New(rand.NewSource(time.Now().UnixNano()))

	referralCode := make([]byte,length)
	for i := range referralCode {
		referralCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(referralCode)
}