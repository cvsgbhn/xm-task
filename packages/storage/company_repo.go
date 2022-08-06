package storage

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"time"
	"xm-task/packages/domain"
	"xm-task/packages/dtl"
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
	sess := r.db.NewSession(nil)

	repComp := dtl.CompanyToDB(c)

	err := sess.InsertInto("companies").
		Returning("id").
		Record(&repComp).
		LoadContext(ctx, &repComp.ID)
	if err != nil {
		return domain.Company{}, err
	}

	return dtl.CompanyFromDB(repComp), nil
}

func (r *CompanyRepo) UpdateCompany(ctx context.Context, c domain.Company) (domain.Company, error) {
	sess := r.db.NewSession(nil)

	repComp := dtl.CompanyToDB(c)

	_, err := sess.Update("companies").
		Where(dbr.Eq("id", repComp.ID)).
		Set("name", repComp.Name).
		Set("country", repComp.Country).
		Set("website", repComp.Website).
		Set("phone", repComp.Phone).
		Set("updated_at", repComp.UpdatedAt).
		ExecContext(ctx)
	if err != nil {
		return c, err
	}

	return dtl.CompanyFromDB(repComp), nil
}

func (r *CompanyRepo) DeleteCompany(ctx context.Context, id int) error {
	sess := r.db.NewSession(nil)

	_, err := sess.Update("companies").
		Where(dbr.Eq("id", id)).
		Set("deleted_at", time.Now().UTC()).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
