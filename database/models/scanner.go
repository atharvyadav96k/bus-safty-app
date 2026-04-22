package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableScanners = "scanners"

	ScannerColID        = "id"
	ScannerColOrgID     = "org_id"
	ScannerColDeviceKey = "device_key"
	ScannerColCreatedAt = "created_at"
	ScannerColUpdatedAt = "updated_at"
	ScannerColDeletedAt = "deleted_at"
)

type Scanner struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrgID     uint           `gorm:"index;not null" json:"org_id"`
	Org       *Org           `gorm:"foreignKey:OrgID;references:ID" json:"-"`
	DeviceKey string         `gorm:"uniqueIndex;not null" json:"device_key"`
	Vehicles  []Vehicle      `gorm:"foreignKey:ScannerID;references:ID" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Scanner) TableName() string {
	return TableScanners
}
