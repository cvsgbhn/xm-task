package domain

import (
	"encoding/json"
	"errors"
	"io"
)

type Storage interface {
	// Companies
	SelectMany() ([]Company, error)
	SelectOne(code string) (Company, error)
	InsertCompany(c Company) (Company, error)
	UpdateCompany(c Company) (Company, error)
	DeleteCompany(id int) error

	//Countries
	SelectCountry(name string) (int, error)
	InsertCountry(name string) (int, error)
}

type CompanyService struct {
	repo Storage
}

func NewCompanyService(r Storage) *CompanyService {
	return &CompanyService{repo: r}
}

func (c *Company) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(r)
}

func (c *Company) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Companies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (s *CompanyService) Create(c Company) (Company, error) {
	cntrID, err := s.repo.SelectCountry(c.Country)
	if errors.Is(err, nil) { // todo dbr error not found
		cntrID, err = s.repo.InsertCountry(c.Country)
	}
	if err != nil || cntrID == 0 {
		return c, nil
	}

	c, err = s.repo.InsertCompany(c)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) Update(code string, c Company) (Company, error) {
	c.Code = code

	c, err := s.repo.UpdateCompany(c)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

// func (s *CompanyService) ShowOne() (Company, error) {}

// func (s *CompanyService) ShowMany() (Companies, error) {}
