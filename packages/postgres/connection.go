package postgres

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	"xm-task/packages/config"
)

func MakeDBconn(cf config.Config) (*dbr.Connection, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cf.Uname,
		cf.Pwd,
		cf.DbName,
		cf.DbHost,
		cf.DbPort,
	)

	db, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
