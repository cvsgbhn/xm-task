package storage

import "time"

type Company struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Country   string    `db:"country"`
	Website   string    `db:"website"`
	Phone     string    `db:"website"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
