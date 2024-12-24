package app

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (a *App) BrowseBorrows() {
	var choice string
	fmt.Print("\033[H\033[2J")
	borrows, err := a.db.GetBorrowsById(a.currentUser.UserId)
	if err != nil {
		a.l.Error("Error getting borrows", "error", err)
		return
	}
	for _, borrow := range borrows {
		fmt.Println(borrow)
	}
	huh.NewSelect[string]().
		Options(huh.NewOptions("back")...).
		Value(&choice).
		Title("Books").Run()
	if choice == "back" {
		a.MainMenu()
	}
}
