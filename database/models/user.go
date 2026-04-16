package database_models

import (
	"time"

	"github.com/atharvyadav96k/gcp/common/entity"
)

type User struct {
	ID                 string          `json:"id" db:"id"`
	Name               string          `json:"name" db:"name"`
	WhiteListedEmailID entity.Email    `json:"white_listed_email_id" db:"white_listed_email_id"`
	ImgURL             string          `json:"img_url" db:"img_url"`
	Password           entity.Password `json:"-" db:"password"`
	OrgID              string          `json:"org_id" db:"org_id"`
	Role               string          `json:"role" db:"role"`
	HasAccess          bool            `json:"has_access" db:"has_access"`
	RFIDID             string          `json:"rfid_id" db:"rfid_id"`
	IsVerified         bool            `json:"is_verified" db:"is_verified"`
	CreatedAt          time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at" db:"update_at"`
}

func main() {
}
