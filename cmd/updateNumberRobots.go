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
	if len(robots) < 1 {
		log.Fatalln("no robots")
	}
	now := time.Now()
	var newRobots []models.RobotsWarehouse
	for i := 0; i < len(robots); i++ {
		robots[i].Count += robots[i].ManufacturingRate
		robots[i].LastUpdateNumberRobots = now
		for k := 0; k < robots[i].ManufacturingRate; k++ {
			newRobots = append(newRobots, models.RobotsWarehouse{
				RobotId:               robots[i].Id,
				Days:                  1,
				StorageCost:           robots[i].StorageCost,
				WarehouseStorageCost:  robots[i].StorageCost,
				LastUpdateStorageCost: now})
		}
	}
	if save := h.DB.Save(&robots); save.Error != nil {
		log.Fatalln(find.Error)
	}
	if save := h.DB.Save(&newRobots); save.Error != nil {
		log.Fatalln(find.Error)
	}
}
