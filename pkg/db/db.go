package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"robot-factory/pkg/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://demo_role:123456@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(models.Robot{})
	//db.AutoMigrate(models.Warehouse{})

	return db
}
