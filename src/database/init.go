package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
}

func NewDBConfig(name string, userName string, password string, host string, port string) *DBConfig {
	return &DBConfig{
		Name:     name,
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (db DBConfig) CreateDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db.Host, db.UserName, db.Password, db.Name, db.Port,
	)
}

func DBConnect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
