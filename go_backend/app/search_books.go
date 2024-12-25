package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) SearchBooks() {
	fmt.Print("\033[H\033[2J")
	var search string
	var bookTitle string
	bookMap := make(map[string]models.Books)
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&search).Title("Search"),

			huh.NewSelect[string]().
				Value(&bookTitle).
				Height(8).
				Title("Books").
				OptionsFunc(func() []huh.Option[string] {
					opts, t := a.db.SearchBooks(search)
					bookMap = t
					return opts
				}, &search),
		),
	).Run()

	a.BookPage(bookMap[bookTitle])
}
