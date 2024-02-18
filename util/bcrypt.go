package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
}

func ComparePassword(hashedPwd []byte, pwd []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPwd, pwd)
}
