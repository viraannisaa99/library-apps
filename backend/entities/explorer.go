package entities

import "time"

// BookExplorer mewakili hasil JOIN authors + books + reviews.
type BookExplorer struct {
	BookID        int        `db:"book_id"        json:"book_id"`
	BookTitle     string     `db:"book_title"     json:"book_title"`
	PublishedYear int        `db:"published_year" json:"published_year"`
	AuthorID      int        `db:"author_id"      json:"author_id"`
	AuthorName    string     `db:"author_name"    json:"author_name"`
	AuthorEmail   string     `db:"author_email"   json:"author_email"`
	ReviewCount   int        `db:"review_count"   json:"review_count"`
	AvgRating     float64    `db:"avg_rating"     json:"avg_rating"`
	LastReviewAt  *time.Time `db:"last_review_at" json:"last_review_at,omitempty"`
}
