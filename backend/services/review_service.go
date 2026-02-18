package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type ReviewService interface {
	GetAll() ([]entities.Review, error)
	GetByID(id int) (*entities.Review, error)
	GetByBookID(bookID int) ([]entities.Review, error)
	Create(req entities.CreateReviewRequest) (*entities.Review, error)
	Update(id int, req entities.UpdateReviewRequest) (*entities.Review, error)
	Delete(id int) error
}

type reviewService struct {
	repo     repositories.ReviewRepository
	bookRepo repositories.BookRepository
}

func NewReviewService(repo repositories.ReviewRepository, bookRepo repositories.BookRepository) ReviewService {
	return &reviewService{repo, bookRepo}
}

func (s *reviewService) GetAll() ([]entities.Review, error) {
	return s.repo.FindAll()
}

func (s *reviewService) GetByID(id int) (*entities.Review, error) {
	review, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("review not found")
		}
		return nil, err
	}
	return review, nil
}

func (s *reviewService) GetByBookID(bookID int) ([]entities.Review, error) {
	if _, err := s.bookRepo.FindByID(bookID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return s.repo.FindByBookID(bookID)
}

func (s *reviewService) Create(req entities.CreateReviewRequest) (*entities.Review, error) {
	if _, err := s.bookRepo.FindByID(req.BookID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return s.repo.Create(req)
}

func (s *reviewService) Update(id int, req entities.UpdateReviewRequest) (*entities.Review, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *reviewService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
