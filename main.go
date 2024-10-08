package main

import (
	"database/sql"
	"log"

	"github.com/blacknoise228/simplebank-backend-learning/api"
	db "github.com/blacknoise228/simplebank-backend-learning/db/sqlc"
	"github.com/blacknoise228/simplebank-backend-learning/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("DB not connected", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("server create error:", err)
	}

	if err = server.Start(config.ServerURL); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
