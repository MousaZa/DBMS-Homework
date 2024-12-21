package app

import "fmt"

func (a *App) BrowseAllBorrows() {
	fmt.Print("\033[H\033[2J")
	borrows, err := a.db.GetBorrows()
	if err != nil {
		a.l.Error("Error getting borrows", "error", err)
		return
	}
	for _, borrow := range borrows {
		fmt.Println(borrow)
	}
	a.MainMenu()
}
