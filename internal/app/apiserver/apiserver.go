package apiserver

import (
	"database/sql"
	"github.com/trad3r/rest-api-go/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseUrl)

	if err != nil {
		return err
	}

	defer db.Close()

	s := sqlstore.New(db)
	srv := NewServer(s)
	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
