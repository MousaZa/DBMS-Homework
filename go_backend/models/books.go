package models

type Books struct {
	BookId     uint64 `json:"bookid" gorm:"column:book_id;primaryKey"`
	Title      string `json:"title" gorm:"column:title"`
	Author     string `json:"author" gorm:"column:author"`
	CategoryId int    `json:"categoryid" gorm:"column:category_id"`
	Language   string `json:"language" gorm:"column:language"`
	Likes      int    `json:"likes" gorm:"column:likes"`
	Summary    string `json:"summary" gorm:"column:summary"`
	Avaliable  bool   `json:"avaliable" gorm:"column:avaliable"`
}
