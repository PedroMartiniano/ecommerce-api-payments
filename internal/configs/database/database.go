package database

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/env"
)

func InitPostgresDatabase() (*sql.DB, error) {
	conn := env.GetEnv("DATABASE_URL")
	DB, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
