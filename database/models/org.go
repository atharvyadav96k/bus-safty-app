package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableOrgs = "orgs"

	OrgColID           = "id"
	OrgColName         = "name"
	OrgColCode         = "code"
	OrgColContactEmail = "contact_email"
	OrgColLogoURL      = "logo_url"
	OrgColRootUserID   = "root_user_id"
	OrgColCreatedAt    = "created_at"
	OrgColUpdatedAt    = "updated_at"
	OrgColDeletedAt    = "deleted_at"
)

type Org struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"not null;index" json:"name"`
	Code         int64          `gorm:"uniqueIndex;not null" json:"code"`
	ContactEmail string         `gorm:"index" json:"contact_email"`
	LogoURL      string         `json:"logo_url"`
	RootUserID   uint           `gorm:"index" json:"root_user_id"`
	RootUser     *RootUser      `gorm:"foreignKey:RootUserID;references:ID" json:"-"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Org) TableName() string {
	return TableOrgs
}
