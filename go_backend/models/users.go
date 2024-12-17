package models

type Users struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Customers struct {
	UserId    int  `json:"user_id"`
	CanBorrow bool `json:"can_borrow"`
}

type Admins struct {
	UserId      int    `json:"user_id"`
	LibraryName string `json:"library_name"`
}
