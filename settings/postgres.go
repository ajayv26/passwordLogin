package settings

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DBClient *sql.DB
)

func ConnectPostgres(dbSource string) (*sql.DB, error) {
	client, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	DBClient = client
	if err = client.Ping(); err != nil {
		return nil, err
	}
	log.Println("successfully connected with postgres db")
	return client, nil
}
