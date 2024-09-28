package user

import (
	"errors"
	"startup/internal/database"
)

func GetAllUsers() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (User, error) {
	var user User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func CreateUser(user *User) error {
	result := database.DB.Create(user)
	return result.Error
}

func UpdateUser(user *User) error {
	result := database.DB.Save(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DB.Delete(&User{}, id)
	return result.Error
}
