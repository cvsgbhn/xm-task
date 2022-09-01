package domain

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"strconv"
	"time"
	"xm-task/packages/entities"
)

type Storage interface {
	// Companies
	SelectMany(ctx context.Context, f entities.Filter) ([]entities.Company, error)
	SelectByID(ctx context.Context, code int64) (entities.Company, error)
	InsertCompany(ctx context.Context, c entities.Company, ctrID int) (entities.Company, error)
	UpdateCompany(ctx context.Context, c entities.Company, ctrID int) (entities.Company, error)
	DeleteCompany(ctx context.Context, id int64) error

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
	if !errors.Is(err, dbr.ErrNotFound) {
		cID, err = s.repo.InsertCountry(ctx, c.Country)
	}
	if err != nil || cID == 0 {
		return c, err
	}

	err = nil

	c.UpdatedAt = time.Now().UTC()

	c, err = s.repo.InsertCompany(ctx, c, cID)
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

	err = nil

	c.Code = code
	c.UpdatedAt = time.Now().UTC()

	c, err = s.repo.UpdateCompany(ctx, c, cID)
	if err != nil || c.Code == "" {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) ShowByCode(ctx context.Context, code string) (entities.Company, error) {
	d, _ := strconv.ParseInt(code, 16, 64)

	c, err := s.repo.SelectByID(ctx, d)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CompanyService) ShowMany(ctx context.Context, f entities.Filter) (entities.Companies, error) {
	c, err := s.repo.SelectMany(ctx, f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CompanyService) DeleteByCode(ctx context.Context, code string) error {
	d, _ := strconv.ParseInt(code, 16, 64)

	err := s.repo.DeleteCompany(ctx, d)
	if err != nil {
		return err
	}

	return nil
}
