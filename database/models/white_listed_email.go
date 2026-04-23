package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableWhiteListedEmails = "white_listed_emails"

	WhiteListedEmailColID        = "id"
	WhiteListedEmailColEmail     = "email"
	WhiteListedEmailColCreatedAt = "created_at"
	WhiteListedEmailColUpdatedAt = "updated_at"
	WhiteListedEmailColDeletedAt = "deleted_at"
)

type WhiteListedEmail struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (WhiteListedEmail) TableName() string {
	return TableWhiteListedEmails
}
