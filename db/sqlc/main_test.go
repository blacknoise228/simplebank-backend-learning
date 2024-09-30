package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "host=localhost port=5432 dbname=simple_bank user=root password=dialogitetatet connect_timeout=10 sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("DB not connected", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
