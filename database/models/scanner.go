package database_models

import "time"

type Scanner struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	DeviceKey string    `json:"device_key" unique="true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Scanner) SetID(id string) {
	s.ID = id
}
