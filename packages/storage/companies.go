package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Company - contains the company data.
type Company struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Country   string    `json:"country"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

type Companies []Company

// CompStorage - simple data storage.
type CompStorage struct {
	Comps Companies
}

func NewStorage() *CompStorage {
	//d := Companies(make([]Company, 0))
	d := Companies{
		Company{
			ID:      1,
			Name:    "qwer",
			Code:    "qwer",
			Country: "qwer",
			Website: "qwer",
		},
		Company{
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

// AddCompany adds data piece to storage and deletes the expired.
func (cs *CompStorage) AddCompany() {
	cs.Comps = append(cs.Comps, Company{
		ID:      3,
		Name:    "a",
		Code:    "b",
		Country: "c",
		Website: "d",
	})
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

func (cs *CompStorage) findCompany(id int64) (*Company, int, error) {
	for i, j := range cs.Comps {
		if j.ID == id {
			return &j, i, nil
		}
	}

	return nil, 0, ErrCompanyNotFound
}

func (cs *CompStorage) UpdateCompany(id int64, cmp *Company) error {
	_, pos, err := cs.findCompany(id)
	if err != nil {
		return err
	}

	cmp.ID = id
	cs.Comps[pos] = *cmp

	return nil
}
