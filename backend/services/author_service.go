package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type AuthorService interface {
	GetAll() ([]entities.Author, error)
	GetByID(id int) (*entities.Author, error)
	Create(req entities.CreateAuthorRequest) (*entities.Author, error)
	Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error)
	Delete(id int) error
}

type authorService struct {
	repo repositories.AuthorRepository
}

func NewAuthorService(repo repositories.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) GetAll() ([]entities.Author, error) {
	return s.repo.FindAll()
}

func (s *authorService) GetByID(id int) (*entities.Author, error) {
	author, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return author, nil
}

func (s *authorService) Create(req entities.CreateAuthorRequest) (*entities.Author, error) {
	return s.repo.Create(req)
}

func (s *authorService) Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error) {
	// Pastikan data ada dulu
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *authorService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
