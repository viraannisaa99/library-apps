package services

import (
	"database/sql"
	"errors"

	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/repositories"
)

type AuthorService struct {
	repo *repositories.AuthorRepository
}

func NewAuthorService(repo *repositories.AuthorRepository) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) GetAll() ([]entities.Author, error) {
	authors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// contoh logic untuk cek apakah author terisi
	if len(authors) == 0 {
		return []entities.Author{}, nil
	}

	return authors, nil
}

func (s *AuthorService) GetByID(id int) (*entities.Author, error) {
	author, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}
	return author, nil
}

func (s *AuthorService) Create(req entities.CreateAuthorRequest) (*entities.Author, error) {
	return s.repo.Create(req)
}

func (s *AuthorService) Update(id int, req entities.UpdateAuthorRequest) (*entities.Author, error) {
	// Pastikan data ada dulu
	if _, err := s.GetByID(id); err != nil {
		return nil, err
	}
	return s.repo.Update(id, req)
}

func (s *AuthorService) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
