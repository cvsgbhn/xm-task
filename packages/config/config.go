package config

import "os"

type Config struct {
	Uname   string
	Pwd     string
	DbName  string
	DbHost  string
	DbPort  string
	AppPort string
}

func GetConfig() Config {
	return Config{
		Uname:   os.Getenv("POSTGRES_USERNAME"),
		Pwd:     os.Getenv("POSTGRES_PWD"),
		DbName:  os.Getenv("POSTGRES_DBNAME"),
		DbHost:  os.Getenv("POSTGRES_HOST"),
		DbPort:  os.Getenv("POSTGRES_PORT"),
		AppPort: os.Getenv("PORT"),
	}
}
