package postgres

import (
	"api-social-network/internal/infrastructure/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB initializes a new database connection using GORM.
func NewDB(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.DBConnectionString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")
	return db, nil
}
