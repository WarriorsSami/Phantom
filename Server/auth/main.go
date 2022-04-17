package main

import (
	"flag"
	"github.com/WarriorsSami/Phantom/Server/db"
	"github.com/joho/godotenv"
	"log"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run service locally")
	flag.Parse()
}

func main() {
	if local {
		if err := godotenv.Load(); err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
}
