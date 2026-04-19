package database_models

import "time"

type RootUser struct {
	ID        int64     `json:"id"`
	Email     int64     `json:"email" unique="true"`
	OrgID     int64     `json:"org_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
