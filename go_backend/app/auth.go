package app

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) Auth() {
	fmt.Print("\033[H\033[2J")
	var choice string

	huh.NewSelect[string]().
		Title("What do you want to do?").
		Options(
			huh.NewOption("Register", "R"),
			huh.NewOption("Login", "L"),
		).
		Value(&choice).Run()

	if choice == "R" {
		a.register()
	} else if choice == "L" {
		a.login()
	}
}

func (a *App) register() {
	var role string
	var name string
	var email string
	var library_name string
	var password string

	register := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose an account type").
				Options(
					huh.NewOption("Customer", "customer"),
					huh.NewOption("Admin", "Admin"),
				).
				Value(&role),
		),

		// Gather some final details about the order.
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Value(&name),

			huh.NewInput().
				Title("Email").
				Value(&email),

			huh.NewInput().
				Title("Password").
				Value(&password),
		),
	)
	err := register.Run()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	if role == "Admin" {
		huh.NewInput().
			Title("Library name").
			Value(&library_name).Run()
	}

	user := &models.Users{}
	var q string
	if role == "Admin" {
		q = fmt.Sprintf("INSERT INTO admins (username, email, password, role, library_name) VALUES ('%s', '%s', '%s', '%s', '%s')", name, email, password, role, library_name)
	} else {
		q = fmt.Sprintf("INSERT INTO customers (username, email, password, role) VALUES ('%s', '%s', '%s', '%s')", name, email, password, role)
	}
	err = a.db.DB.Exec(q, user).Error
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	a.Auth()
}

func (a *App) login() {
	var email string
	var password string

	register := huh.NewForm(

		huh.NewGroup(
			huh.NewInput().
				Title("Email").
				Value(&email),

			huh.NewInput().
				Title("Password").
				Value(&password),
		),
	)

	err := register.Run()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	user := &models.Users{}

	q := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' AND password = '%s'", email, password)

	err = a.db.DB.Raw(q).Scan(&user).Error
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	if user.UserId == 0 {
		fmt.Println("User not found")
		a.Auth()
	}
	a.currentUser = user

	if user.Role == "Admin" {
		a.AdminMainMenu()
	} else {
		a.UserMainMenu()
	}

}
