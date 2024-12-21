package models

type Categories struct {
	CategoryId   int    `gorm:"column:category_id"`
	CategoryName string `gorm:"column:name"`
}
