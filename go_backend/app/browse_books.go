package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) BrowseBooks() {
	fmt.Print("\033[H\033[2J")
	var bookName string
	books, err := a.db.GetBooks()
	if err != nil {
		a.l.Error("Error getting books", "error", err)
		return
	}
	bookTitles := make([]huh.Option[string], len(books))
	bookMap := make(map[string]models.Books, len(books))
	for i, book := range books {
		bookTitles[i] = huh.NewOption(book.Title, book.Title)
		bookMap[book.Title] = book
	}

	huh.NewSelect[string]().
		Options(bookTitles...).
		Value(&bookName).
		Title("Books").Run()

	selectedBook := bookMap[bookName]

	a.BookPage(selectedBook)
}
