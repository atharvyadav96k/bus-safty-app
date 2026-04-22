package database_models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableUsers = "users"

	UserColID                 = "id"
	UserColName               = "name"
	UserColWhiteListedEmailID = "white_listed_email_id"
	UserColImgURL             = "img_url"
	UserColPassword           = "password"
	UserColOrgID              = "org_id"
	UserColRoleID             = "role_id"
	UserColHasAccess          = "has_access"
	UserColRFIDID             = "rfid_id"
	UserColIsVerified         = "is_verified"
	UserColCreatedAt          = "created_at"
	UserColUpdatedAt          = "updated_at"
	UserColDeletedAt          = "deleted_at"
)

type User struct {
	ID                 uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	Name               string            `gorm:"not null;index" json:"name"`
	WhiteListedEmailID uint              `gorm:"index" json:"white_listed_email_id"`
	WhiteListedEmail   *WhiteListedEmail `gorm:"foreignKey:WhiteListedEmailID;references:ID" json:"-"`
	ImgURL             string            `json:"img_url"`
	Password           string            `json:"password"`
	OrgID              uint              `gorm:"index;not null" json:"org_id"`
	Org                *Org              `gorm:"foreignKey:OrgID;references:ID" json:"-"`
	RoleID             uint              `gorm:"index" json:"role_id"`
	HasAccess          bool              `gorm:"default:false" json:"has_access"`
	RFIDID             uint              `gorm:"index" json:"rfid_id"`
	RFID               *RFID             `gorm:"foreignKey:RFIDID;references:ID" json:"-"`
	IsVerified         bool              `gorm:"default:false" json:"is_verified"`
	CreatedAt          time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt    `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return TableUsers
}
