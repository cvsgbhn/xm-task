package storage

import (
	"github.com/gocraft/dbr/v2"
)

type Repo struct {
	*CompanyRepo
	*CountryRepo
}

type dbConn struct {
	*dbr.Connection
}

func NewRepository(db *dbr.Connection) *Repo {
	base := &dbConn{db}

	return &Repo{
		CompanyRepo: NewCompanyRepo(base),
		CountryRepo: NewCountryRepo(base),
	}
}
