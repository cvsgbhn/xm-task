package storage

import (
	"context"
	"github.com/gocraft/dbr/v2"
)

type CountryRepo struct {
	sess dbr.Session
	db   dbConn
}

func NewCountryRepo(db *dbConn) *CountryRepo {
	return &CountryRepo{db: *db}
}

func (r *CountryRepo) SelectCountryID(ctx context.Context, name string) (int, error) {
	sess := r.db.NewSession(nil)

	var cID int

	err := sess.Select("id").
		From("countries").
		Where(dbr.Eq("name", name)).
		LoadOne(&cID)
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (r *CountryRepo) InsertCountry(ctx context.Context, name string) (int, error) {
	sess := r.db.NewSession(nil)

	c := &Country{Name: name}

	err := sess.InsertInto("countries").
		Returning("id").
		Pair("name", c.Name).
		LoadContext(ctx, &c.ID)
	if err != nil {
		return 0, err
	}

	return c.ID, nil
}
