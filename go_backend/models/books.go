package models

type Books struct {
	BookId     uint64 `json:"bookid" gorm:"column:book_id;primaryKey"`
	Title      string `json:"title" gorm:"title"`
	Author     string `json:"author" gorm:"author"`
	CategoryId int    `json:"categoryid" gorm:"categoryid"`
	Language   string `json:"language" gorm:"language"`
	Likes      int    `json:"likes" gorm:"likes"`
	Summary    string `json:"summary" gorm:"summary"`
	Avaliable  bool   `json:"avaliable" gorm:"avaliable"`
}
