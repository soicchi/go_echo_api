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
