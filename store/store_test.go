package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost user=trad3r password=psql dbname=go_rest_api_test sslmode=disable"
	}

	os.Exit(m.Run())
}
