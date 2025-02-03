package utils

import (
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", dto.ErrHashPass
	}

	return string(hashPass), nil
}
