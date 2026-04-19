package database_models

import (
	"time"

	"github.com/atharvyadav96k/gcp/common/entity"
)

type Org struct {
	ID           string       `json:"id" firestore:"id"`
	Name         string       `json:"name" firestore:"name"`
	Code         string       `json:"code" firestore:"code" unique:"true"`
	ContactEmail entity.Email `json:"contact_email" firestore:"contact_email"`
	LogoURL      string       `json:"logo_url" firestore:"logo_url"`
	RootUserID   string       `json:"root_user_id" firestore:"root_user_id"`
	SubStart     time.Time    `json:"sub_start" firestore:"sub_start"`
	SubEnd       time.Time    `json:"sub_end" firestore:"sub_end"`
	CreatedAt    time.Time    `json:"created_at" firestore:"create_at"`
	UpdatedAt    time.Time    `json:"updated_at" firestore:"update_at"`
}

func (o *Org) SetID(id string) {
	o.ID = id
}
