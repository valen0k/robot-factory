package main

import (
	"log"
	"robot-factory/pkg/db"
	"robot-factory/pkg/handlers"
	"robot-factory/pkg/models"
	"robot-factory/pkg/util"
	"time"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	DB, err := db.Init(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	h := handlers.New(DB)

	var robots []models.Robot
	find := h.DB.Where("count_of_robots > 0 AND last_update_storage_cost < CURRENT_DATE").Find(&robots)
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
			CountRobots:       robots[i].CountOfRobots,
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
