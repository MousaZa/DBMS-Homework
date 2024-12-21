package models

type Likes struct {
	LikeId int `gorm:"column:like_id"`
	UserId int `gorm:"column:user_id"`
	BookId int `gorm:"column:book_id"`
}
