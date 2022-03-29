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
	if len(robots) < 1 {
		log.Fatalln("no robots")
	}

	var sale []models.TransactionHistory
	now := time.Now()

	for i := 0; i < len(robots); i++ {
		sale = append(sale, models.TransactionHistory{
			Transaction:       models.STORAGE,
			RobotId:           robots[i].Id,
			CountRobots:       robots[i].Count,
			Amount:            robots[i].StorageCost,
			ManufacturingCost: 0,
			Time:              now,
		})
		robots[i].LastUpdateStorageCost = now
	}

	if save1 := h.DB.Save(&robots); save1.Error != nil {
		log.Fatalln(save1.Error)
	}

	if save2 := h.DB.Save(&sale); save2.Error != nil {
		log.Fatalln(save2.Error)
	}
}
