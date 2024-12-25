package app

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (a *App) BrowseNotifications() {
	fmt.Print("\033[H\033[2J")
	notifications, err := a.db.GetNotifications(a.currentUser.UserId)
	if err != nil {
		a.l.Error("Error getting notifications", "error", err)
		return
	}
	// bookTitles := make([]huh.Option[string], len(notifications))
	var n string
	for i, notification := range notifications {
		n = n + fmt.Sprintf("%d. %s\n", i+1, notification.Message)
		// bookMap[book.Title] = book
	}
	var choice string
	huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Description(n).
				Title("Notifications"),
			huh.NewSelect[string]().
				Options(huh.NewOptions("back")...).
				Value(&choice).
				Title("Select"),
		),
	).Run()
	if choice == "back" {
		a.MainMenu()
	}
}
