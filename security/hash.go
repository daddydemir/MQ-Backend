package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(data string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(pass)
}

func CheckPassword(encyripted string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encyripted), []byte(password))
	if err == nil {
		return true
	}
	return false
}
