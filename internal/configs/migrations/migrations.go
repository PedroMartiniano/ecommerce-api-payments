package migrations

import (
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigrations() error {
	migrationsPath := env.GetEnv("MIGRATIONS_PATH")
	dbUrl := env.GetEnv("DATABASE_URL")

	m, err := migrate.New(
		migrationsPath,
		dbUrl,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
