package domain

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"strconv"
	"strings"
	"time"
	"xm-task/packages/entities"
)

type Storage interface {
	// Companies
	SelectMany(ctx context.Context, f entities.Filter) ([]entities.Company, error)
	SelectByID(ctx context.Context, code int64) (entities.Company, error)
	InsertCompany(ctx context.Context, c entities.Company) (entities.Company, error)
	UpdateCompany(ctx context.Context, c entities.Company) (entities.Company, error)
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

func (s *CompanyService) Create(ctx context.Context, c entities.Company) (entities.Company, error) {
	cID, err := s.repo.SelectCountryID(ctx, c.Country)
	if errors.Is(err, dbr.ErrNotFound) {
		cID, err = s.repo.InsertCountry(ctx, c.Country)
	}
	if err != nil || cID == 0 {
		return c, err
	}

	c.UpdatedAt = time.Now().UTC()

	c, err = s.repo.InsertCompany(ctx, c)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) Update(ctx context.Context, code string, c entities.Company) (entities.Company, error) {
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

func (s *CompanyService) ShowByCode(ctx context.Context, code string) (entities.Company, error) {
	dStr := strings.Replace(code, "0x", "", -1)
	dStr = strings.Replace(dStr, "0X", "", -1)
	d, _ := strconv.ParseInt(dStr, 16, 64)

	c, err := s.repo.SelectByID(ctx, d)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) ShowMany(ctx context.Context, f entities.Filter) (entities.Companies, error) {
	return nil, nil
}

func (s *CompanyService) DeleteByCode(ctx context.Context, code string) error {
	return nil
}
