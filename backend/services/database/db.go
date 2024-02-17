package database

import (
	"habitTrackerApi/services/domains"

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

func Migrate(db *gorm.DB) error {
	// Migrate Models
	if err := db.AutoMigrate(&domains.UserClient{}, &domains.Habit{}); err != nil {
		return err
	}

	return nil
}
