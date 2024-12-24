package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/MousaZa/DBMS-Homework/go_backend/storage"
	"github.com/hashicorp/go-hclog"
)

type App struct {
	currentUser *models.Users
	db          *storage.Database
	l           hclog.Logger
}

func NewApp(db *storage.Database) *App {
	return &App{
		l:  hclog.Default(),
		db: db,
	}
}

func (a *App) StartApp() {
	fmt.Print("\033[H\033[2J")
	a.Auth()
}
