package controllers

import (
	"errors"
	"hitos_back/models"
	"hitos_back/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := models.User{}

	err = models.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	// if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// 	return "", err
	// }
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
func GetUserByID(uid uint) (models.User, error) {

	var u models.User

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
