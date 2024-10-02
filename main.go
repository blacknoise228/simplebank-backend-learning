package main

import (
	"database/sql"
	"log"

	"github.com/blacknoise228/simplebank-backend-learning/api"
	db "github.com/blacknoise228/simplebank-backend-learning/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "host=localhost port=5432 dbname=simple_bank user=root password=password connect_timeout=10 sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("DB not connected", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
