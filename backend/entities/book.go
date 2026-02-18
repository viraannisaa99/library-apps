package entities

import "time"

type Book struct {
	ID            int       `db:"id"             json:"id"`
	AuthorID      int       `db:"author_id"      json:"author_id"`
	Title         string    `db:"title"          json:"title"`
	Description   string    `db:"description"    json:"description"`
	PublishedYear int       `db:"published_year" json:"published_year"`
	CreatedAt     time.Time `db:"created_at"     json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"     json:"updated_at"`

	// Joined field (tidak disimpan di db)
	Author *Author `db:"-" json:"author,omitempty"`
}

type CreateBookRequest struct {
	AuthorID      int    `json:"author_id"      binding:"required"`
	Title         string `json:"title"          binding:"required"`
	Description   string `json:"description"`
	PublishedYear int    `json:"published_year"`
}

type UpdateBookRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	PublishedYear int    `json:"published_year"`
}
