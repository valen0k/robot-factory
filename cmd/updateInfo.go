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
	find := h.DB.Where("").Find(&robots)
	if find.Error != nil {
		log.Fatalln(find.Error)
	}
	now := time.Now()
	for i := 0; i < len(robots); i++ {
		if robots[i].LastUpdate.Day() != now.Day() {
			if robots[i].Count > 0 {
				robots[i].Allowance += robots[i].StorageCost
			} else {
				robots[i].Allowance = 0
			}
			robots[i].Count += robots[i].ManufacturingRate
			robots[i].LastUpdate = now
		}
	}
	if save := h.DB.Save(&robots); save.Error != nil {
		log.Fatalln(find.Error)
	}
}
