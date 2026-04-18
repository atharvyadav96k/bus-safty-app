package database_models

import (
	"time"

	"github.com/atharvyadav96k/gcp/common/entity"
)

type Org struct {
	ID           string       `json:"id" db:"id"`
	Name         string       `json:"name" db:"name"`
	Code         string       `json:"code" db:"code" unique:"true"`
	ContactEmail entity.Email `json:"contact_email" db:"contact_email"`
	LogoURL      string       `json:"logo_url" db:"logo_url"`
	RootUserID   string       `json:"root_user_id" db:"root_user_id"`
	SubStart     time.Time    `json:"sub_start" db:"sub_start"`
	SubEnd       time.Time    `json:"sub_end" db:"sub_end"`
	CreatedAt    time.Time    `json:"created_at" db:"create_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"update_at"`
}
