package db

import (
	"github.com/luis501angel/go-course/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	user, exist, _ := CheckIfUserExist(email)
	if !exist {
		return user, false
	}

	passwordBytes := []byte(password)
	userPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(userPassword, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true
}
