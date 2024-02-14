package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func ConnectToDb() *sqlx.DB {
	conn, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatalf("❌Failed to connect to database: %v❌", err)
	}

	err = goose.SetDialect(os.Getenv("DB_DRIVER"))
	if err != nil {
		log.Fatalf("❌Failed to set dialect: %v❌", err)
	}

	err = goose.Up(conn.DB, os.Getenv("DB_MIGRATIONS"))
	if err != nil {
		log.Fatalf("❌Failed to apply migrations: %v❌", err)
	}

	fmt.Println("✅Migrations applied successfully.✅")
	return conn
}
