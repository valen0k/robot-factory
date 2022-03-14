package handlers

import (
	"log"
	"robot-factory/pkg/models"
	"time"
)

const triggerTime = "23:42:"

func (h handler) Trigger() {
	for tick := range time.Tick(time.Minute) {
		//if strings.Contains(tick.String(), triggerTime) {
		// Prints UTC time and date
		log.Println(tick)
		var robots []models.Robot
		if find := h.DB.Find(&robots); find.Error != nil {
			log.Println(find.Error)
		} else {
			for i := 0; i < len(robots); i++ {
				if robots[i].Count > 0 {
					robots[i].Allowance += robots[i].StorageCost
				} else {
					robots[i].Allowance = 0
				}
				robots[i].Count += robots[i].ManufacturingRate
			}
			if save := h.DB.Save(&robots); save.Error != nil {
				log.Println(find.Error)
			}
		}
		//}
	}
}
