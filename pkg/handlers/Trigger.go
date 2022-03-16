package handlers

import (
	"log"
	"robot-factory/pkg/models"
)

func (h handler) Trigger() {
	log.Println("tut1")
	var robots []models.Robot
	find := h.DB.Find(&robots)
	if find.Error != nil {
		log.Fatalln(find.Error)
		return
	}
	for i := 0; i < len(robots); i++ {
		if robots[i].Count > 0 {
			robots[i].Allowance += robots[i].StorageCost
		} else {
			robots[i].Allowance = 0
		}
		robots[i].Count += robots[i].ManufacturingRate
	}
	if save := h.DB.Save(&robots); save.Error != nil {
		log.Fatalln(find.Error)
		return
	}
	log.Println("tut2")
}
