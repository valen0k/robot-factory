package models

import "time"

type transaction string

const (
	STORAGE transaction = "STORAGE"
	SALE                = "SALE"
)

type TransactionHistory struct {
	Id                int         `gorm:"primaryKey"`
	Transaction       transaction `gorm:"type:varchar"`
	RobotId           int
	CountRobots       int
	Amount            int
	ManufacturingCost int
	Time              time.Time
}
