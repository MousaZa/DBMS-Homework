package models

type Book struct {
	BookId     int    `json:"book_id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CategoryId int    `json:"category_id"`
	CoverUrl   string `json:"cover_url"`
	Language   string `json:"language"`
	Likes      int    `json:"likes"`
	Summary    string `json:"summary"`
	Avaliable  bool   `json:"avaliable"`
}
