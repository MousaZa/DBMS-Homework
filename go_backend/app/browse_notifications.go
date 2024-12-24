package app

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (a *App) BrowseNotifications() {
	fmt.Print("\033[H\033[2J")
	var bookName string
	books, err := a.db.GetNotifications(a.currentUser.UserId)
	if err != nil {
		a.l.Error("Error getting books", "error", err)
		return
	}
	bookTitles := make([]huh.Option[string], len(books))
	for i, book := range books {
		bookTitles[i] = huh.NewOption(book.Message, book.Message)
		// bookMap[book.Title] = book
	}

	huh.NewSelect[string]().
		Options(bookTitles...).
		Value(&bookName).
		Title("Notifications").Run()

	a.MainMenu()
}
