package models

import "time"

type Sale struct {
	Id          int `gorm:"primaryKey"`
	RobotId     int
	CountRobots int
	SellTime    time.Time
	Profit      int
}
