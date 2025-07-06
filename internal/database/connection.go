package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/projects-for-public/help-govern/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB

func Connect() error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL environment variable not set")
	}
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func ConnectGORM(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
}
