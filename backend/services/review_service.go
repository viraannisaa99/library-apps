package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type ReviewService struct {
	repo     *repositories.ReviewRepository
	bookRepo *repositories.BookRepository
}

func NewReviewService(repo *repositories.ReviewRepository, bookRepo *repositories.BookRepository) *ReviewService {
	return &ReviewService{repo, bookRepo}
}

func (s *ReviewService) GetAll() ([]entities.Review, error) {
	return s.repo.GetAll()
}

func (s *ReviewService) GetByID(id int) (*entities.Review, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrReviewNotFound
		}
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) GetByBookID(bookID int) ([]entities.Review, error) {
	if _, err := s.bookRepo.GetByID(bookID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return s.repo.GetByBookID(bookID)
}

func (s *ReviewService) Create(req entities.CreateReviewRequest) (*entities.Review, error) {
	if _, err := s.bookRepo.GetByID(req.BookID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return s.repo.Create(req)
}

func (s *ReviewService) Update(id int, req entities.UpdateReviewRequest) (*entities.Review, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *ReviewService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
