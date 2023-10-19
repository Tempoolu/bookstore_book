package models

import "time"

type Book struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
}

func (b *Book) First() {
	db.First(b)
}

func (b *Book) Create(title, content string) {
	b.Title = title
	b.Content = content
	_ = db.Create(b)
}
