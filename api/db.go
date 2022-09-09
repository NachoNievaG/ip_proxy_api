package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	dbname = "ip2location"
)

func ConnectToDB() *sql.DB {
	password := os.Getenv("PSQL_PASS")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("could not connect to psql %s", err)
	}
	return result
}
