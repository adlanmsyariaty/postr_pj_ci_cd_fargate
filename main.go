package main

import (
	"database/sql"
	"log"

	"github.com/adlanmsyariaty/postr/api"
	"github.com/adlanmsyariaty/postr/config"
	db "github.com/adlanmsyariaty/postr/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create new server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
