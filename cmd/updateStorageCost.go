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
	find := h.DB.Where("last_update_storage_cost < CURRENT_DATE").Find(&robots)
	if find.Error != nil {
		log.Fatalln(find.Error)
	}
	if len(robots) > 0 {
		now := time.Now()
		for i := 0; i < len(robots); i++ {
			if robots[i].Count != 0 {
				robots[i].WarehouseStorageCost += robots[i].StorageCost
			} else {
				robots[i].WarehouseStorageCost = 0
			}
			robots[i].LastUpdateStorageCost = now
		}
		if save := h.DB.Save(&robots); save.Error != nil {
			log.Fatalln(find.Error)
		}
	}
}
