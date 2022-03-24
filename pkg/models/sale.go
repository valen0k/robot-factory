package models

import "time"

type Sale struct {
	Id                int `gorm:"primaryKey"`
	RobotId           int
	CountRobots       int
	ManufacturingCost int
	SellingPrice      int
	SellTime          time.Time
}
