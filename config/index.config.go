package appConfig

import (
	"os"
)

func DefineApiPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	res := ":" + port
	return res
}

func DefineDatabaseDSN() string {
	var dsn string
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbName == "" {
		dbName = "postgres"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	dsn = "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName
	return dsn
}
