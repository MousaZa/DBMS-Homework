package models

type Users struct {
	UserId   uint64 `json:"user_id" gorm:"column:user_id;primaryKey"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
}

type Customers struct {
	UserId    uint64 `json:"user_id"`
	CanBorrow bool   `json:"can_borrow"`
}

type Admins struct {
	UserId      uint64 `json:"user_id"`
	LibraryName string `json:"library_name"`
}
