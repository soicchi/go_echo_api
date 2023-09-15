package database

import (
	"fmt"

	"go_echo_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

func NewDBConfig(name string, userName string, password string, host string, port string, sslMode string) *DBConfig {
	return &DBConfig{
		Name:     name,
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
		SSLMode:  sslMode,
	}
}

func (db DBConfig) CreateDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.Host, db.UserName, db.Password, db.Name, db.Port, db.SSLMode,
	)
}

func DBConnect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
