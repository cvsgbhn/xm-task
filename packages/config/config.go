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
		Uname:   os.Getenv("DB_USER"),
		Pwd:     os.Getenv("DB_PWD"),
		DbName:  os.Getenv("DB_NAME"),
		DbHost:  os.Getenv("DB_HOST"),
		DbPort:  os.Getenv("DB_PORT"),
		AppPort: os.Getenv("PORT"),
	}
}
