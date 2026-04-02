package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type BookService struct {
	repo       *repositories.BookRepository
	authorRepo *repositories.AuthorRepository
}

func NewBookService(repo *repositories.BookRepository, authorRepo *repositories.AuthorRepository) *BookService {
	return &BookService{repo, authorRepo}
}

func (s *BookService) GetAll() ([]entities.Book, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetByID(id int) (*entities.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return book, nil
}

func (s *BookService) GetByAuthorID(authorID int) ([]entities.Book, error) {
	// Validasi author ada
	if _, err := s.authorRepo.GetByID(authorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}
	return s.repo.GetByAuthorID(authorID)
}

func (s *BookService) GetExplorer(authorID int, minRating float64) ([]entities.BookExplorer, error) {
	if authorID > 0 {
		if _, err := s.authorRepo.GetByID(authorID); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrAuthorNotFound
			}
			return nil, err
		}
	}

	if minRating < 0 || minRating > 5 {
		return nil, ErrInvalidMinRating
	}

	return s.repo.GetExplorer(authorID, minRating)
}

func (s *BookService) Create(req entities.CreateBookRequest) (*entities.Book, error) {
	// Validasi author_id valid
	if _, err := s.authorRepo.GetByID(req.AuthorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}
	return s.repo.Create(req)
}

func (s *BookService) Update(id int, req entities.UpdateBookRequest) (*entities.Book, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *BookService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
