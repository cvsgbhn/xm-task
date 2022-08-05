package storage

import (
	"context"
	"xm-task/packages/domain"
)

type CompanyRepo struct {
	db dbConn
}

func NewCompanyRepo(db *dbConn) *CompanyRepo {
	return &CompanyRepo{db: *db}
}

func (r *CompanyRepo) SelectMany(ctx context.Context) ([]domain.Company, error) {
	return nil, nil
}

func (r *CompanyRepo) SelectOne(ctx context.Context, code string) (domain.Company, error) {
	return domain.Company{}, nil
}

func (r *CompanyRepo) InsertCompany(ctx context.Context, c domain.Company) (domain.Company, error) {
	return domain.Company{}, nil
}

func (r *CompanyRepo) UpdateCompany(ctx context.Context, c domain.Company) (domain.Company, error) {
	return domain.Company{}, nil
}

func (r *CompanyRepo) DeleteCompany(ctx context.Context, id int) error {
	return nil
}
