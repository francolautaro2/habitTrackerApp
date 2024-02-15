package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Contect to database
func ConnectDb() (*gorm.DB, error) {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/habittracker?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
