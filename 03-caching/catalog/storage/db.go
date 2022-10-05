package storage

import (
	"database/sql"
	"fmt"
	"log"
)

type DBDriver interface {
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

const (
	dbCredentialsFormat = "user=%s password=%s dbname=%s host=%s port=%d sslmode=%s"
)

func NewDB(cfg DBConfig) *sql.DB {
	address := fmt.Sprintf(dbCredentialsFormat,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Host,
		cfg.Port,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", address)

	if err != nil {
		log.Fatal("[Database] failed connecting to DB: " + address + ", err: " + err.Error())
	}

	// check established connection with DB
	if err := db.Ping(); err != nil {
		log.Fatal("[Database] db is unreachable: " + address + ", err: " + err.Error())
	}

	return db
}
