package app

import (
	"fmt"
	"time"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (a *App) BookPage(book models.Books) {
	fmt.Print("\033[H\033[2J")
	var choice string
	fmt.Println("Id: ", book.BookId)
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author)
	fmt.Println("Summary: ", book.Summary)
	fmt.Println("Language: ", book.Language)
	fmt.Println("Likes: ", book.Likes)

	options := huh.NewOptions("like", "borrow", "back")

	if a.currentUser.Role == "Admin" {
		options = append(options, huh.NewOption("delete", "delete"))
	}

	huh.NewSelect[string]().
		Options(options...).
		Value(&choice).
		Title("Books").Run()
	if choice == "like" {
		a.LikeBook(book.BookId)
	} else if choice == "borrow" {
		a.db.BorrowBook(a.currentUser, book)
	} else if choice == "delete" {
		err := a.db.DeleteBook(book.BookId)
		if err != nil {
			a.l.Error("Error deleting book", "error", err)
			return
		}
		a.MainMenu()
	} else {
		a.MainMenu()
	}

}

func (a *App) LikeBook(b uint64) {
	a.l.Info("Liking book", "book", b, "user", a.currentUser.UserId)

	q := fmt.Sprintf("INSERT INTO likes (user_id, book_id) VALUES (%v, %v)", a.currentUser.UserId, b)
	err := a.db.DB.Exec(q).Error
	if err != nil {
		a.l.Error("Error liking book", "error", err)
		return
	}
	fmt.Print("\033[H\033[2J")
	time.Sleep(1 * time.Second)
	a.MainMenu()
}
