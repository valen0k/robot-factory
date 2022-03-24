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
	var updateInfo []models.RobotsWarehouse
	find := h.DB.Where("last_update_storage_cost < CURRENT_DATE AND sale_id > 0").Find(&updateInfo)
	if find.Error != nil {
		log.Fatalln(find.Error)
	}
	if len(updateInfo) < 1 {
		log.Fatalln("all update")
	}
	now := time.Now()
	for i := 0; i < len(updateInfo); i++ {
		updateInfo[i].Days++
		updateInfo[i].WarehouseStorageCost += updateInfo[i].StorageCost
		updateInfo[i].LastUpdateStorageCost = now
	}
	if save := h.DB.Save(&updateInfo); save.Error != nil {
		log.Fatalln(find.Error)
	}
}
