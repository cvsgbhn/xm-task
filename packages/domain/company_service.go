package domain

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gocraft/dbr/v2"
	"io"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	// Companies
	SelectMany(ctx context.Context) ([]Company, error)
	SelectByID(ctx context.Context, code int64) (Company, error)
	InsertCompany(ctx context.Context, c Company) (Company, error)
	UpdateCompany(ctx context.Context, c Company) (Company, error)
	DeleteCompany(ctx context.Context, id int) error

	//Countries
	SelectCountryID(ctx context.Context, name string) (int, error)
	InsertCountry(ctx context.Context, name string) (int, error)
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

func (s *CompanyService) Create(ctx context.Context, c Company) (Company, error) {
	cID, err := s.repo.SelectCountryID(ctx, c.Country)
	if errors.Is(err, dbr.ErrNotFound) {
		cID, err = s.repo.InsertCountry(ctx, c.Country)
	}
	if err != nil || cID == 0 {
		return c, nil
	}

	c.UpdatedAt = time.Now().UTC()

	c, err = s.repo.InsertCompany(ctx, c)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) Update(ctx context.Context, code string, c Company) (Company, error) {
	cID, err := s.repo.SelectCountryID(ctx, c.Country)
	if errors.Is(err, dbr.ErrNotFound) {
		cID, err = s.repo.InsertCountry(ctx, c.Country)
	}
	if err != nil || cID == 0 {
		return c, nil
	}

	c.Code = code
	c.UpdatedAt = time.Now().UTC()

	c, err = s.repo.UpdateCompany(ctx, c)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) ShowByCode(ctx context.Context, code string) (Company, error) {
	dStr := strings.Replace(code, "0x", "", -1)
	dStr = strings.Replace(dStr, "0X", "", -1)
	d, _ := strconv.ParseInt(dStr, 16, 64)

	c, err := s.repo.SelectByID(ctx, d)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) ShowMany(ctx context.Context, f Filter) (Companies, error) {
	return nil, nil
}

func (s *CompanyService) DeleteByCode(ctx context.Context, code string) error {
	return nil
}
