package domain

import "time"

type Company struct {
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Country   string    `json:"country"`
	Website   string    `json:"website"`
	Phone     string    `json:"website"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Companies []Company
