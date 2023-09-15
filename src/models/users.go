package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u User) Create(db *gorm.DB) error {
	result := db.Create(&u)
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}

func (u User) UpdateUser(db *gorm.DB, userID string) error {
	err := db.Model(&u).Where("id = ?", userID).Update("name", u.Name)
	if err.Error != nil {
		return fmt.Errorf("failed to update user: %w", err.Error)
	}

	return nil
}

func DeleteUser(db *gorm.DB, userID string) error {
	var u User
	err := db.Where("id = ?", userID).Delete(&u)
	if err.Error != nil {
		return fmt.Errorf("failed to delete user: %w", err.Error)
	}

	return nil
}
