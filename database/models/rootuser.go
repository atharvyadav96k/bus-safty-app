package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableRootUsers = "root_users"

	RootUserColID        = "id"
	RootUserColEmail     = "email"
	RootUserColOrgID     = "org_id"
	RootUserColPassword  = "password"
	RootUserColCreatedAt = "created_at"
	RootUserColUpdatedAt = "updated_at"
	RootUserColDeletedAt = "deleted_at"
)

type RootUser struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	OrgID     uint           `gorm:"index" json:"org_id"`
	Org       *Org           `gorm:"foreignKey:OrgID;references:ID" json:"-"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RootUser) TableName() string {
	return TableRootUsers
}
