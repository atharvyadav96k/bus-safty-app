package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableRFIDs = "rfids"

	RFIDColID         = "id"
	RFIDColRFIDNumber = "rfid_number"
	RFIDColOrgID      = "org_id"
	RFIDColCreatedAt  = "created_at"
	RFIDColUpdatedAt  = "updated_at"
	RFIDColDeletedAt  = "deleted_at"
)

type RFID struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	RFIDNumber string         `gorm:"uniqueIndex;not null" json:"rfid_number"`
	OrgID      uint           `gorm:"index;not null" json:"org_id"`
	Org        *Org           `gorm:"foreignKey:OrgID;references:ID" json:"-"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RFID) TableName() string {
	return TableRFIDs
}
