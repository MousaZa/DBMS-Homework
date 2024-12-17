package models

type Likes struct {
	LikeId int `json:"like_id"`
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
