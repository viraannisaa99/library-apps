package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vira/go-crud/entities"
)

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{db}
}

func (r *AuthorRepository) GetAll() ([]entities.Author, error) {
	authors := []entities.Author{}
	query := `SELECT * FROM authors ORDER BY id`
	err := r.db.Select(&authors, query)
	return authors, err
}

func (r *AuthorRepository) GetByID(id int) (*entities.Author, error) {
	author := &entities.Author{}
	query := `SELECT * FROM authors WHERE id = $1`
	err := r.db.Get(author, query, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (r *AuthorRepository) Create(req entities.CreateAuthorRequest) (*entities.Author, error) {
	author := &entities.Author{}
	query := `
		INSERT INTO authors (name, email, bio)
		VALUES ($1, $2, $3)
		RETURNING *`
	err := r.db.Get(author, query, req.Name, req.Email, req.Bio)
	return author, err
}

func (r *AuthorRepository) Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error) {
	author := &entities.Author{}
	query := `
		UPDATE authors
		SET name       = COALESCE(NULLIF($1, ''), name),
		    email      = COALESCE(NULLIF($2, ''), email),
		    bio        = COALESCE(NULLIF($3, ''), bio),
		    updated_at = NOW()
		WHERE id = $4
		RETURNING *`
	err := r.db.Get(author, query, req.Name, req.Email, req.Bio, id)
	return author, err
}

func (r *AuthorRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM authors WHERE id = $1`, id)
	return err
}
