package models

import "time"

type RobotsWarehouse struct {
	Id                    int `gorm:"primaryKey"`
	RobotId               int
	SaleId                int
	Days                  int
	StorageCost           int
	WarehouseStorageCost  int
	LastUpdateStorageCost time.Time
}
