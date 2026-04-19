package database_models

import "time"

type Vehicle struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	BusNumber string    `json:"bus_number" unique:"true"`
	ScannerID int64     `json:"scanner_id"`
	Lat       string    `json:"lat"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v *Vehicle) SetID(id string) {
	v.ID = id
}
