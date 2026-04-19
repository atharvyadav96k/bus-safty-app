package database_models

import (
	"time"

	"github.com/atharvyadav96k/gcp/common/entity"
)

type User struct {
	ID                 string          `json:"id" firestore:"id"`
	Name               string          `json:"name" firestore:"name"`
	WhiteListedEmailID entity.Email    `json:"white_listed_email_id" firestore:"white_listed_email_id" unique:"true"`
	ImgURL             string          `json:"img_url" firestore:"img_url"`
	Password           entity.Password `json:"password" firestore:"password"`
	OrgID              string          `json:"org_id" firestore:"org_id"`
	Role               string          `json:"role" firestore:"role"`
	HasAccess          bool            `json:"has_access" firestore:"has_access"`
	RFIDID             string          `json:"rfid_id" firestore:"rfid_id"`
	IsVerified         bool            `json:"is_verified" firestore:"is_verified"`
	CreatedAt          time.Time       `json:"created_at" firestore:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at" firestore:"update_at"`
}

func (u *User) SetID(id string) {
	u.ID = id
}
