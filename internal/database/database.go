package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/cronui/cronui/internal/entity"
)

func NewDatabase(gl logger.Interface) (*gorm.DB, error) {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		log.Println("cannot create data directory")
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{
		Logger: gl,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
