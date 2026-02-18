package entities

import "time"

type Review struct {
	ID        int       `db:"id"         json:"id"`
	BookID    int       `db:"book_id"    json:"book_id"`
	Reviewer  string    `db:"reviewer"   json:"reviewer"`
	Rating    int       `db:"rating"     json:"rating"`
	Comment   string    `db:"comment"    json:"comment"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	// Joined field
	Book *Book `db:"-" json:"book,omitempty"`
}

type CreateReviewRequest struct {
	BookID   int    `json:"book_id"  binding:"required"`
	Reviewer string `json:"reviewer" binding:"required"`
	Rating   int    `json:"rating"   binding:"required,min=1,max=5"`
	Comment  string `json:"comment"`
}

type UpdateReviewRequest struct {
	Reviewer string `json:"reviewer"`
	Rating   int    `json:"rating"  binding:"omitempty,min=1,max=5"`
	Comment  string `json:"comment"`
}
