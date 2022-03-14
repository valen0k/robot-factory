package db

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"robot-factory/pkg/models"
)

const dbURL = "postgres://demo_role:123456@localhost:5432/postgres"

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err == nil {
		if err := db.AutoMigrate(models.Robot{}); err != nil {
			return nil, errors.New(err.Error())
		} else if err := db.AutoMigrate(models.Sale{}); err != nil {
			return nil, errors.New(err.Error())
		} else {
			return db, nil
		}
	}
	return nil, errors.New(err.Error())
}
