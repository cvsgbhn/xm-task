package storage

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"time"
	"xm-task/packages/dbmodels"
	"xm-task/packages/dtl"
	"xm-task/packages/entities"
)

type CompanyRepo struct {
	db dbConn
}

func NewCompanyRepo(db *dbConn) *CompanyRepo {
	return &CompanyRepo{db: *db}
}

func (r *CompanyRepo) SelectMany(ctx context.Context, f entities.Filter) ([]entities.Company, error) {
	sess := r.db.NewSession(nil)

	c := make([]dbmodels.Company, 0)

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

	_, err := stmt.Where(dbr.Eq("deleted_at", nil)).LoadContext(ctx, &c)
	if err != nil {
		return nil, err
	}

	return dtl.CompaniesFromDB(c), nil
}

func (r *CompanyRepo) SelectByID(ctx context.Context, code int64) (entities.Company, error) {
	sess := r.db.NewSession(nil)

	c := dbmodels.Company{}

	err := sess.Select("companies.id", "companies.name", "countries.name as country", "website", "phone", "updated_at").
		From("companies").
		Join("countries", "countries.id = companies.country").
		Where(dbr.Eq("companies.id", code)).
		Where(dbr.Eq("deleted_at", nil)).
		LoadOneContext(ctx, &c)
	if err != nil {
		return entities.Company{}, err
	}

	return dtl.CompanyFromDB(c), err
}

func (r *CompanyRepo) InsertCompany(ctx context.Context, c entities.Company, ctrID int) (entities.Company, error) {
	sess := r.db.NewSession(nil)

	repComp := dtl.CompanyToDB(c)
	repComp.CreatedAt = time.Now().UTC()
	repComp.UpdatedAt = repComp.CreatedAt

	err := sess.InsertInto("companies").
		Returning("id").
		Pair("name", repComp.Name).
		Pair("country", ctrID).
		Pair("website", repComp.Website).
		Pair("phone", repComp.Phone).
		Pair("created_at", repComp.CreatedAt).
		Pair("updated_at", repComp.UpdatedAt).
		LoadContext(ctx, &repComp.ID)
	if err != nil {
		return entities.Company{}, err
	}

	return dtl.CompanyFromDB(repComp), nil
}

func (r *CompanyRepo) UpdateCompany(ctx context.Context, c entities.Company, ctrID int) (entities.Company, error) {
	sess := r.db.NewSession(nil)

	repComp := dtl.CompanyToDB(c)

	_, err := sess.Update("companies").
		Where(dbr.Eq("id", repComp.ID)).
		Set("name", repComp.Name).
		Set("updated_at", repComp.UpdatedAt).
		Set("country", ctrID).
		Set("website", repComp.Website).
		Set("phone", repComp.Phone).
		ExecContext(ctx)
	if err != nil {
		return c, err
	}

	return dtl.CompanyFromDB(repComp), nil
}

func (r *CompanyRepo) DeleteCompany(ctx context.Context, id int64) error {
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
