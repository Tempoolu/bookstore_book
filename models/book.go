package models

import "time"

type Book struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
}

func (b *Book) List() (books []Book) {
	db.Find(&books)
	return books
}

func (b *Book) Get(id string) {
	db.First(b, id)
}

func (b *Book) Create() {
	_ = db.Create(b)
}

func (b *Book) Update(id string) {
	db.Model(b).Where("id = ?", id).Updates(map[string]interface{}{"title": b.Title, "content": b.Content})
}
