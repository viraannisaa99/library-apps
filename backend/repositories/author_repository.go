package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vira/go-crud/entities"
)

type AuthorRepository interface {
	FindAll() ([]entities.Author, error)
	FindByID(id int) (*entities.Author, error)
	Create(req entities.CreateAuthorRequest) (*entities.Author, error)
	Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error)
	Delete(id int) error
}

type authorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll() ([]entities.Author, error) {
	authors := []entities.Author{}
	query := `SELECT * FROM authors ORDER BY id`
	err := r.db.Select(&authors, query)
	return authors, err
}

func (r *authorRepository) FindByID(id int) (*entities.Author, error) {
	author := &entities.Author{}
	query := `SELECT * FROM authors WHERE id = $1`
	err := r.db.Get(author, query, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (r *authorRepository) Create(req entities.CreateAuthorRequest) (*entities.Author, error) {
	author := &entities.Author{}
	query := `
		INSERT INTO authors (name, email, bio)
		VALUES ($1, $2, $3)
		RETURNING *`
	err := r.db.Get(author, query, req.Name, req.Email, req.Bio)
	return author, err
}

func (r *authorRepository) Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error) {
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

func (r *authorRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM authors WHERE id = $1`, id)
	return err
}
