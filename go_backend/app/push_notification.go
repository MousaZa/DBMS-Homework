package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) PushNotification() {
	fmt.Print("\033[H\033[2J")
	var text string
	var usr string
	users, err := a.db.GetUsers()
	if err != nil {
		a.l.Error("Error getting books", "error", err)
		return
	}
	usersTitles := make([]huh.Option[string], len(users))
	userMap := make(map[string]models.Users, len(users))
	for i, user := range users {
		usersTitles[i] = huh.NewOption(user.Username, user.Username)
		userMap[user.Username] = user
	}
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Text").
				Value(&text),

			huh.NewSelect[string]().
				Options(usersTitles...).
				Value(&usr).
				Title("User"),
		),
	).Run()

	err = a.db.PushNotification(text, userMap[usr].UserId)
	if err != nil {
		a.l.Error("Error adding book", "error", err)
		return
	}
	a.MainMenu()
}
