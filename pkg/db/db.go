package db

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"robot-factory/pkg/models"
	"robot-factory/pkg/util"
)

const dbURL = "postgres://demo_role:123456@localhost:5432/postgres"

func Init(config util.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if err := db.AutoMigrate(models.Robot{}, models.TransactionHistory{}); err != nil {
		return nil, errors.New(err.Error())
	}

	return db, nil
}
