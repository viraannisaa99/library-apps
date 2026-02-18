package entities

import "time"

type Author struct {
	ID        int       `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"        binding:"required"`
	Email     string    `db:"email"      json:"email"       binding:"required,email"`
	Bio       string    `db:"bio"        json:"bio"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Request DTOs
type CreateAuthorRequest struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Bio   string `json:"bio"`
}

type UpdateAuthorRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
	Bio   string `json:"bio"`
}
