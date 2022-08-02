package storage

type CountryRepo struct {
	db dbConn
}

func NewCountryRepo(db *dbConn) *CountryRepo {
	return &CountryRepo{db: *db}
}

func (r *CountryRepo) SelectCountry(name string) (int, error) {}

func (r *CountryRepo) InsertCountry(name string) (int, error) {}
