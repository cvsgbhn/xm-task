package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Company - contains the company data.
type CompanyDraft struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Country   string    `json:"country"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

type Companies []CompanyDraft

// CompStorage - simple data storage.
type CompStorage struct {
	Comps Companies
}

func NewStorage() *CompStorage {
	//d := Companies(make([]Company, 0))
	d := Companies{
		CompanyDraft{
			ID:      1,
			Name:    "qwer",
			Code:    "qwer",
			Country: "qwer",
			Website: "qwer",
		},
		CompanyDraft{
			ID:      2,
			Name:    "asf",
			Code:    "asdf",
			Country: "asfd",
			Website: "asf",
		},
	}

	return &CompStorage{
		Comps: d,
	}
}

func (c *Company) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(r)
}

func (ad *Companies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ad)
}

// GetCompanies shows all data from storage.
func (cs *CompStorage) GetCompanies() Companies {
	return cs.Comps
}

var ErrCompanyNotFound = fmt.Errorf("Company not found")
