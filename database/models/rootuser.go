package database_models

import "time"

type RootUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" unique="true"`
	OrgID     string    `json:"org_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *RootUser) SetID(id string) {
	r.ID = id
}
