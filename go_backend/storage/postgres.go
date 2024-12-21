package storage

import (
	"fmt"

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

func NewConnection(cfg *Config) (*Database, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port) // Adjust TimeZone as needed

	config := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
		l:  hclog.Default(),
	}, nil

}

type Database struct {
	DB *gorm.DB
	l  hclog.Logger
}
