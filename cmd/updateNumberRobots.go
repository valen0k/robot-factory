package main

import (
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
	h := handlers.New(DB)
	var robots []models.Robot
	find := h.DB.Where("last_update_number_robots < CURRENT_DATE").Find(&robots)
	if find.Error != nil {
		log.Fatalln(find.Error)
	}
	if len(robots) > 0 {
		now := time.Now()
		for i := 0; i < len(robots); i++ {
			robots[i].Count += robots[i].ManufacturingRate
			robots[i].LastUpdateNumberRobots = now
		}
		if save := h.DB.Save(&robots); save.Error != nil {
			log.Fatalln(find.Error)
		}
	}
}
