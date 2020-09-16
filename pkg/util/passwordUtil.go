package util

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func MatchPassword(password string, encodePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
