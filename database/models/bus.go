package database_models

import "time"

type Vehicle struct {
	ID        int64     `json:"id"`
	OrgID     int64     `json:"org_id"`
	BusNumber string    `json:"bus_number" unique:"true"`
	ScannerID int64     `json:"scanner_id"`
	Lat       string    `json:"lat"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
