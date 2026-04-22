package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableVehicles = "vehicles"

	VehicleColID        = "id"
	VehicleColOrgID     = "org_id"
	VehicleColBusNumber = "bus_number"
	VehicleColScannerID = "scanner_id"
	VehicleColLat       = "lat"
	VehicleColLong      = "long"
	VehicleColCreatedAt = "created_at"
	VehicleColUpdatedAt = "updated_at"
	VehicleColDeletedAt = "deleted_at"
)

type Vehicle struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrgID     uint           `gorm:"index;not null" json:"org_id"`
	Org       *Org           `gorm:"foreignKey:OrgID;references:ID" json:"-"`
	BusNumber string         `gorm:"index;not null" json:"bus_number"`
	ScannerID uint           `gorm:"index" json:"scanner_id"`
	Scanner   *Scanner       `gorm:"foreignKey:ScannerID;references:ID" json:"-"`
	Lat       string         `json:"lat"`
	Long      string         `json:"long"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Vehicle) TableName() string {
	return TableVehicles
}
