package storage

import (
	"fmt"

	"github.com/MousaZa/DBMS-Homework/go_backend/models"
	"github.com/charmbracelet/huh"
)

func (DB *Database) PushNotification(text string, id uint64) error {
	return DB.DB.Exec("INSERT INTO commercial_notifications (message, user_id, status) VALUES (?, ?, ?)", text, id, "sent").Error

}

func (DB *Database) SearchBooks(search string) []huh.Option[string] {

	var books []models.Books
	q := fmt.Sprintf("SELECT * FROM books WHERE title LIKE '%%%s%%'", search)
	err := DB.DB.Raw(q).Scan(&books).Error
	if err != nil {
		panic(err)
	}
	bookTitles := make([]huh.Option[string], len(books))
	for i, book := range books {
		bookTitles[i] = huh.NewOption(book.Title, book.Title)
	}
	return bookTitles
}

func (DB *Database) GetUsers() ([]models.Users, error) {
	var users []models.Users
	err := DB.DB.Raw("SELECT * FROM users").Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (DB *Database) GetNotifications(id uint64) ([]models.CommercialNotifications, error) {
	var not []models.CommercialNotifications
	err := DB.DB.Raw("SELECT * FROM commercial_notifications WHERE user_id = ?", id).Scan(&not).Error
	if err != nil {
		return nil, err
	}
	return not, nil
}

func (DB *Database) GetBooks() ([]models.Books, error) {
	var books []models.Books
	err := DB.DB.Raw("SELECT * FROM books").Scan(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (DB *Database) GetBorrows() ([]models.Borrows, error) {
	var borrows []models.Borrows
	err := DB.DB.Raw("SELECT * FROM borrows").Scan(&borrows).Error
	if err != nil {
		return nil, err
	}
	return borrows, nil

}
func (DB *Database) GetBorrowsById(id uint64) ([]models.Borrows, error) {
	var borrows []models.Borrows
	err := DB.DB.Raw("SELECT * FROM borrows WHERE user_id = ?", id).Scan(&borrows).Error
	if err != nil {
		return nil, err
	}
	return borrows, nil

}

func (DB *Database) GetCategories() ([]models.Categories, error) {
	var categories []models.Categories
	err := DB.DB.Raw("SELECT * FROM categories").Scan(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (DB *Database) DeleteBook(id uint64) error {
	err := DB.DB.Exec("DELETE FROM books WHERE book_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (DB *Database) AddBook(book models.Books) error {
	err := DB.DB.Exec("INSERT INTO books (title, author, summary, language, category_id) VALUES (?, ?, ?, ?, ?)", book.Title, book.Author, book.Summary, book.Language, book.CategoryId).Error
	if err != nil {
		return err
	}
	return nil
}

func (DB *Database) BorrowBook(*models.Users, models.Books) {

}
