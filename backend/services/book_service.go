package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type BookService interface {
	GetAll() ([]entities.Book, error)
	GetByID(id int) (*entities.Book, error)
	GetByAuthorID(authorID int) ([]entities.Book, error)
	GetExplorer(authorID int, minRating float64) ([]entities.BookExplorer, error)
	Create(req entities.CreateBookRequest) (*entities.Book, error)
	Update(id int, req entities.UpdateBookRequest) (*entities.Book, error)
	Delete(id int) error
}

type bookService struct {
	repo       repositories.BookRepository
	authorRepo repositories.AuthorRepository
}

func NewBookService(repo repositories.BookRepository, authorRepo repositories.AuthorRepository) BookService {
	return &bookService{repo, authorRepo}
}

func (s *bookService) GetAll() ([]entities.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetByID(id int) (*entities.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return book, nil
}

func (s *bookService) GetByAuthorID(authorID int) ([]entities.Book, error) {
	// Validasi author ada
	if _, err := s.authorRepo.FindByID(authorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return s.repo.FindByAuthorID(authorID)
}

func (s *bookService) GetExplorer(authorID int, minRating float64) ([]entities.BookExplorer, error) {
	if authorID > 0 {
		if _, err := s.authorRepo.FindByID(authorID); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.New("author not found")
			}
			return nil, err
		}
	}

	if minRating < 0 || minRating > 5 {
		return nil, errors.New("min_rating must be between 0 and 5")
	}

	return s.repo.FindExplorer(authorID, minRating)
}

func (s *bookService) Create(req entities.CreateBookRequest) (*entities.Book, error) {
	// Validasi author_id valid
	if _, err := s.authorRepo.FindByID(req.AuthorID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return s.repo.Create(req)
}

func (s *bookService) Update(id int, req entities.UpdateBookRequest) (*entities.Book, error) {
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *bookService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
