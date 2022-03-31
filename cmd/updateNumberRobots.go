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
	where := h.DB.Model(models.Robot{}).
		Where("last_update_number_robots < CURRENT_DATE").
		Update("count_of_robots", gorm.Expr("count_of_robots + manufacturing_rate")).
		Update("last_update_number_robots", now)
	if where.Error != nil {
		log.Fatalln(where.Error)
	}
	//var robots []models.Robot
	//find := h.DB.Where("last_update_number_robots < CURRENT_DATE").Find(&robots)
	//if find.Error != nil {
	//	log.Fatalln(find.Error)
	//}
	//if len(robots) < 1 {
	//	log.Fatalln("no robots")
	//}
	//for i := 0; i < len(robots); i++ {
	//	robots[i].CountOfRobots += robots[i].ManufacturingRate
	//	robots[i].LastUpdateNumberRobots = now
	//}
	//if save := h.DB.Save(&robots); save.Error != nil {
	//	log.Fatalln(find.Error)
	//}
}
