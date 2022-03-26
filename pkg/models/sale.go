package models

import "time"

type Sale struct {
	Id          int `gorm:"primaryKey"`
	Transaction string
	RobotId     int
	CountRobots int
	Cost        int
	SellPrice   int
	SellTime    time.Time
}
