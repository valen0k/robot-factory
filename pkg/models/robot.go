package models

import "time"

type Robot struct {
	Id                     int       `json:"id" gorm:"primaryKey"`
	Type                   string    `json:"type" gorm:"unique"` // robot type, 3 or more characters
	CountOfRobots          int       `json:"count_of_robots"`
	ManufacturingCost      int       `json:"manufacturing_cost"` // coins per unit
	StorageCost            int       `json:"storage_cost"`       // coins per unit per day
	SellingPrice           int       `json:"selling_price"`      // coins per unit
	ManufacturingRate      int       `json:"manufacturing_rate"` // units per day
	LastUpdateNumberRobots time.Time `json:"-"`
	LastUpdateStorageCost  time.Time `json:"-"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
