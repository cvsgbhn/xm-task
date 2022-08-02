package storage

import "xm-task/packages/domain"

type CompanyRepo struct {
	db dbConn
}

func NewCompanyRepo(db *dbConn) *CompanyRepo {
	return &CompanyRepo{db: *db}
}

func (r *CompanyRepo) SelectMany() ([]domain.Company, error) {
	return nil, nil
}

func (r *CompanyRepo) SelectOne(code string) (domain.Company, error) {
	return domain.Company{}, nil
}

func (r *CompanyRepo) InsertCompany(c domain.Company) (domain.Company, error) {}

func (r *CompanyRepo) UpdateCompany(c domain.Company) (domain.Company, error) {}

func (r *CompanyRepo) DeleteCompany(id int) error {}
