package storage

import (
	"fmt"

	// "github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConnection(cfg *Config) (*gorm.DB, error) {

	// bin, err := ioutil.ReadFile("/run/secrets/db-password")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read database password: %w", err)
	// }

	// Build connection string with SSL mode disabled
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port) // Adjust TimeZone as needed

	// Open connection using GORM
	config := &gorm.Config{}
	return gorm.Open(postgres.Open(dsn), config)

}

type Database struct {
	db *gorm.DB
	l  hclog.Logger
}

// func (db *Database) Migrate() {
// 	err := db.db.AutoMigrate(&models.Borrow{})
// 	if err != nil {
// 		db.l.Error("Unable to migrate borrows", "error", err)
// 		return
// 	}
// 	db.l.Info("borrows migrated successfully")

// }
