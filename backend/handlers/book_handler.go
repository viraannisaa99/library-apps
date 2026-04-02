package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/response"
	"github.com/vira/go-crud/services"
	"github.com/vira/go-crud/utils"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service}
}

// GET /books
func (h *BookHandler) GetAll(c *gin.Context) {
	// Support filter by author: GET /books?author_id=1
	if authorIDStr := c.Query("author_id"); authorIDStr != "" {
		authorID, err := strconv.Atoi(authorIDStr)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "invalid author_id")
			return
		}
		books, err := h.service.GetByAuthorID(authorID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, services.ErrAuthorNotFound) {
				status = http.StatusNotFound
			}
			response.RespondError(c, status, err.Error())
			return
		}
		response.RespondSuccess(c, http.StatusOK, books)
		return
	}

	books, err := h.service.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, books)
}

// GET /books/explorer
func (h *BookHandler) GetExplorer(c *gin.Context) {
	authorID := 0
	minRating := 0.0

	if authorIDStr := c.Query("author_id"); authorIDStr != "" {
		parsedAuthorID, err := strconv.Atoi(authorIDStr)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "invalid author_id")
			return
		}
		authorID = parsedAuthorID
	}

	if minRatingStr := c.Query("min_rating"); minRatingStr != "" {
		parsedMinRating, err := strconv.ParseFloat(minRatingStr, 64)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "invalid min_rating")
			return
		}
		minRating = parsedMinRating
	}

	items, err := h.service.GetExplorer(authorID, minRating)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, services.ErrAuthorNotFound):
			status = http.StatusNotFound
		case errors.Is(err, services.ErrInvalidMinRating):
			status = http.StatusBadRequest
		}
		response.RespondError(c, status, err.Error())
		return
	}

	response.RespondSuccess(c, http.StatusOK, items)
}

// GET /books/:id
func (h *BookHandler) GetByID(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	book, err := h.service.GetByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrBookNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, book)
}

// POST /books
func (h *BookHandler) Create(c *gin.Context) {
	var req entities.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.Create(req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrAuthorNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusCreated, book)
}

// PUT /books/:id
func (h *BookHandler) Update(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	var req entities.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.Update(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrBookNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, book)
}

// DELETE /books/:id
func (h *BookHandler) Delete(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrBookNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, "book deleted")
}
