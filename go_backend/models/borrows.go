package models

type Borrows struct {
	BorrowId  int `json:"borrow_id"`
	UserId    int `json:"user_id"`
	BookId    int `json:"book_id"`
	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`
}
