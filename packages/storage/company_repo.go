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

func (r *CompanyRepo) SelectMany(ctx context.Context, f domain.Filter) ([]domain.Company, error) {
	sess := r.db.NewSession(nil)

	c := make([]Company, 0)

	stmt := sess.Select("id", "name", "country", "website", "phone", "updated_at").From("companies")

	if len(f.Country) > 0 {
		stmt = stmt.Join("countries", dbr.Eq("countries.name", f.Country))
	}

	if len(f.Name) > 0 {
		stmt = stmt.Where(dbr.Eq("name", f.Name))
	}

	if len(f.Phone) > 0 {
		stmt = stmt.Where(dbr.Eq("phone", f.Phone))
	}

	if len(f.Website) > 0 {
		stmt = stmt.Where(dbr.Eq("website", f.Website))
	}

	_, err := stmt.LoadContext(ctx, &c)
	if err != nil {
		return nil, err
	}

	return dtl.CompaniesFromDB(c), nil
}

func (r *CompanyRepo) SelectByID(ctx context.Context, code int64) (domain.Company, error) {
	sess := r.db.NewSession(nil)

	c := Company{}

	err := sess.Select("id", "name", "country", "website", "phone", "updated_at").
		From("companies").
		Where(dbr.Eq("id", code)).
		LoadOneContext(ctx, c)
	if err != nil {
		return domain.Company{}, err
	}

	return dtl.CompanyFromDB(c), err
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
