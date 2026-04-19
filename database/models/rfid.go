package database_models

import "time"

type RFID struct {
	ID         int64     `json:"id"`
	RFIDNumber string    `json:"rfid_number" unique:"true"`
	OrgID      int64     `json:"org_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
