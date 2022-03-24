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
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if err := db.AutoMigrate(models.Robot{}); err != nil {
		return nil, errors.New(err.Error())
	}
	if err := db.AutoMigrate(models.Sale{}); err != nil {
		return nil, errors.New(err.Error())
	}
	if err := db.AutoMigrate(models.RobotsWarehouse{}); err != nil {
		return nil, errors.New(err.Error())
	}
	return db, nil
}
