package app

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (a *App) MainMenu() {
	fmt.Print("\033[H\033[2J")

	if a.currentUser.Role == "Admin" {
		a.AdminMainMenu()
	} else {
		a.UserMainMenu()
	}

}

func (a *App) AdminMainMenu() {
	var choice string
	huh.NewSelect[string]().
		Title("What do you want to do?").
		Options(
			huh.NewOption("Browse books", "browsebooks"),
			huh.NewOption("Browse borrows", "browseborrows"),
			huh.NewOption("Add book", "addbook"),
		).
		Value(&choice).Run()
	if choice == "browsebooks" {
		a.BrowseBooks()
	}
	if choice == "browseborrows" {
		a.BrowseAllBorrows()
	}
	if choice == "addbook" {
		a.AddBook()
	}

}

func (a *App) UserMainMenu() {
	huh.NewSelect[string]().
		Title("What do you want to do?").
		Options(
			huh.NewOption("Browse books", "browse books"),
			huh.NewOption("Browse borrows", "browse borrows"),
		).
		Value(&a.page).Run()
}
