package main

import (
	"gorm.io/gorm"
	"log"
	"robot-factory/pkg/db"
	"robot-factory/pkg/handlers"
	"robot-factory/pkg/models"
	"time"
)

func main() {
	DB, err := db.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}
	now := time.Now()
	h := handlers.New(DB)
	h.DB.Model(models.Robot{}).
		Where("last_update_number_robots < CURRENT_DATE").
		Update("count_of_robots", gorm.Expr("count_of_robots + manufacturing_rate")).
		Update("last_update_number_robots", now)
}
