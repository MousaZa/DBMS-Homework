package models

import "time"

type Borrows struct {
	BorrowId  int       `json:"borrow_id"`
	UserId    int       `json:"user_id"`
	BookId    uint64    `json:"book_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
