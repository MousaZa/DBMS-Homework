package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) EditBook(b models.Books) {
	fmt.Print("\033[H\033[2J")
	var Cat string
	categories, err := a.db.GetCategories()
	if err != nil {
		a.l.Error("Error getting books", "error", err)
		return
	}
	categoriesTitles := make([]huh.Option[string], len(categories))
	bookMap := make(map[string]models.Categories, len(categories))
	for i, category := range categories {
		categoriesTitles[i] = huh.NewOption(category.CategoryName, category.CategoryName)
		bookMap[category.CategoryName] = category
	}

	var book models.Books = b
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Title").
				Value(&book.Title),

			huh.NewInput().
				Title("Author").
				Value(&book.Author),

			huh.NewInput().
				Title("Summary").
				Value(&book.Summary),

			huh.NewSelect[string]().
				Options(huh.NewOptions("Arabic", "English", "Turkish")...).
				Value(&book.Language).
				Title("Language"),

			huh.NewSelect[string]().
				Options(categoriesTitles...).
				Value(&Cat).
				Title("Category"),
		),
	).Run()

	book.CategoryId = bookMap[Cat].CategoryId
	err = a.db.EditBook(book)
	if err != nil {
		a.l.Error("Error adding book", "error", err)
		return
	}
	a.MainMenu()
}
