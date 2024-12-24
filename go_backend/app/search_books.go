package app

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (a *App) SearchBooks() {
	fmt.Print("\033[H\033[2J")
	var search string
	var bookTitle string
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&search).Title("Search"),

			huh.NewSelect[string]().
				Value(&bookTitle).
				Height(8).
				Title("Books").
				OptionsFunc(func() []huh.Option[string] {
					opts := a.db.SearchBooks(search)
					return opts
				}, &search),
		),
	).Run()

	// a.MainMenu()
}
