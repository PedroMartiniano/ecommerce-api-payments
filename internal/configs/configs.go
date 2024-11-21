package configs

import (
	"database/sql"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/database"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/logger"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/migrations"
	"github.com/joho/godotenv"
)

var (
	DB     *sql.DB
	Logger = logger.NewLogger()
)

func InitConfig() {
	var err error

	err = godotenv.Load()
	if err != nil {
		Logger.Debugf("error loading .env file, err: %s", err.Error())
		panic("Error loading .env file")
	}

	DB, err = database.InitPostgresDatabase()
	if err != nil {
		Logger.Errorf("error connecting to postgres database: %s", err.Error())
		panic("Error connecting to postgres database")
	}

	err = migrations.InitMigrations()
	if err != nil {
		Logger.Errorf("error initializing migrations: %s", err.Error())
		panic("Error initializing migrations")
	}
}
