package database

import (
	"projects/config"
	"projects/middlewares"
	"projects/models"

	"github.com/labstack/echo"
)

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user interface{}) (interface{}, error) {
	if err := config.DB.Find(&user, "id=?", id).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
