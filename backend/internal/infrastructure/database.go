// internal/infrastructure/database.go
package database

import (
	"fmt" 
	"log" 
	"github.com/saku-730/web-specimen/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"          
)

//connect Database
func NewDatabaseConnection(cfg *configs.Config) (*gorm.DB, error) {
	dsn := cfg.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = db.Debug()

	if err != nil {
		log.Printf("Failed database connect: %v\n", err)
		return nil, fmt.Errorf("Database connect error: %w", err)
	}

	log.Println("Success connect database")

	return db, nil
}
