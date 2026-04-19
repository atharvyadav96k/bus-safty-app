package database_models

import "time"

type WhiteListedEmail struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" unique="true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *WhiteListedEmail) SetID(id string) {
	w.ID = id
}
