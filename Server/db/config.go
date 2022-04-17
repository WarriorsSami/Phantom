package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	GetDsn() string
	GetDbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg = config{
		dbUser: os.Getenv("DATABASE_USER"),
		dbPass: os.Getenv("DATABASE_PASS"),
		dbHost: os.Getenv("DATABASE_HOST"),
		dbName: os.Getenv("DATABASE_NAME"),
	}

	var err error
	if cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT")); err != nil {
		log.Fatalf("[ERROR] Unable to read port number: %v", err.Error())
	}
	cfg.dsn = fmt.Sprintf("monbodb://%s:%s@%s:%d/%s",
		cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)

	return &cfg
}

func (c *config) GetDsn() string {
	return c.dsn
}

func (c *config) GetDbName() string {
	return c.dbName
}
