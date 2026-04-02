package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vira/go-crud/entities"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetAll() ([]entities.Book, error) {
	books := []entities.Book{}
	query := `SELECT * FROM books ORDER BY id`
	err := r.db.Select(&books, query)
	return books, err
}

func (r *BookRepository) GetByID(id int) (*entities.Book, error) {
	book := &entities.Book{}
	query := `SELECT * FROM books WHERE id = $1`
	err := r.db.Get(book, query, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) GetByAuthorID(authorID int) ([]entities.Book, error) {
	books := []entities.Book{}
	query := `SELECT * FROM books WHERE author_id = $1 ORDER BY id`
	err := r.db.Select(&books, query, authorID)
	return books, err
}

func (r *BookRepository) GetExplorer(authorID int, minRating float64) ([]entities.BookExplorer, error) {
	items := []entities.BookExplorer{}
	query := `
		SELECT
			b.id AS book_id,
			b.title AS book_title,
			b.published_year,
			b.author_id,
			a.name AS author_name,
			a.email AS author_email,
			COUNT(r.id) AS review_count,
			COALESCE(AVG(r.rating)::float8, 0) AS avg_rating,
			MAX(r.created_at) AS last_review_at
		FROM books b
		JOIN authors a ON a.id = b.author_id
		LEFT JOIN reviews r ON r.book_id = b.id
		WHERE ($1 = 0 OR b.author_id = $1)
		GROUP BY b.id, b.title, b.published_year, b.author_id, a.name, a.email
		HAVING ($2 = 0 OR COALESCE(AVG(r.rating)::float8, 0) >= $2)
		ORDER BY b.id`
	err := r.db.Select(&items, query, authorID, minRating)
	return items, err
}

func (r *BookRepository) Create(req entities.CreateBookRequest) (*entities.Book, error) {
	book := &entities.Book{}
	query := `
		INSERT INTO books (author_id, title, description, published_year)
		VALUES ($1, $2, $3, $4)
		RETURNING *`
	err := r.db.Get(book, query, req.AuthorID, req.Title, req.Description, req.PublishedYear)
	return book, err
}

func (r *BookRepository) Update(id int, req entities.UpdateBookRequest) (*entities.Book, error) {
	book := &entities.Book{}
	query := `
		UPDATE books
		SET title          = COALESCE(NULLIF($1, ''), title),
		    description    = COALESCE(NULLIF($2, ''), description),
		    published_year = CASE WHEN $3 = 0 THEN published_year ELSE $3 END,
		    updated_at     = NOW()
		WHERE id = $4
		RETURNING *`
	err := r.db.Get(book, query, req.Title, req.Description, req.PublishedYear, id)
	return book, err
}

func (r *BookRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM books WHERE id = $1`, id)
	return err
}
