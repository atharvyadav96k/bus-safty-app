package database_models

import "time"

type Scanner struct {
	ID        int64     `json:"id"`
	OrgID     int64     `json:"org_id"`
	DeviceKey string    `json:"device_key" unique="true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
