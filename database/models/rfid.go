package database_models

import "time"

type RFID struct {
	ID         string    `json:"id"`
	RFIDNumber string    `json:"rfid_number" unique:"true"`
	OrgID      string    `json:"org_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (r *RFID) SetID(id string) {
	r.ID = id
}
