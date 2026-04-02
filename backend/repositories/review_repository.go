package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vira/go-crud/entities"
)

type ReviewRepository struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (r *ReviewRepository) GetAll() ([]entities.Review, error) {
	reviews := []entities.Review{}
	query := `SELECT * FROM reviews ORDER BY id`
	err := r.db.Select(&reviews, query)
	return reviews, err
}

func (r *ReviewRepository) GetByID(id int) (*entities.Review, error) {
	review := &entities.Review{}
	query := `SELECT * FROM reviews WHERE id = $1`
	err := r.db.Get(review, query, id)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewRepository) GetByBookID(bookID int) ([]entities.Review, error) {
	reviews := []entities.Review{}
	query := `SELECT * FROM reviews WHERE book_id = $1 ORDER BY id`
	err := r.db.Select(&reviews, query, bookID)
	return reviews, err
}

func (r *ReviewRepository) Create(req entities.CreateReviewRequest) (*entities.Review, error) {
	review := &entities.Review{}
	query := `
		INSERT INTO reviews (book_id, reviewer, rating, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING *`
	err := r.db.Get(review, query, req.BookID, req.Reviewer, req.Rating, req.Comment)
	return review, err
}

func (r *ReviewRepository) Update(id int, req entities.UpdateReviewRequest) (*entities.Review, error) {
	review := &entities.Review{}
	query := `
		UPDATE reviews
		SET reviewer   = COALESCE(NULLIF($1, ''), reviewer),
		    rating     = CASE WHEN $2 = 0 THEN rating ELSE $2 END,
		    comment    = COALESCE(NULLIF($3, ''), comment),
		    updated_at = NOW()
		WHERE id = $4
		RETURNING *`
	err := r.db.Get(review, query, req.Reviewer, req.Rating, req.Comment, id)
	return review, err
}

func (r *ReviewRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM reviews WHERE id = $1`, id)
	return err
}
