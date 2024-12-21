package storage

import (
	"github.com/MousaZa/DBMS-Homework/go_backend/models"
)

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
