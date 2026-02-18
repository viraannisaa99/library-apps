package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/services"
)

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service}
}

// GET /books
func (h *BookHandler) GetAll(c *gin.Context) {
	// Support filter by author: GET /books?author_id=1
	if authorIDStr := c.Query("author_id"); authorIDStr != "" {
		authorID, err := strconv.Atoi(authorIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author_id"})
			return
		}
		books, err := h.service.GetByAuthorID(authorID)
		if err != nil {
			status := http.StatusInternalServerError
			if err.Error() == "author not found" {
				status = http.StatusNotFound
			}
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": books})
		return
	}

	books, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/explorer
func (h *BookHandler) GetExplorer(c *gin.Context) {
	authorID := 0
	minRating := 0.0

	if authorIDStr := c.Query("author_id"); authorIDStr != "" {
		parsedAuthorID, err := strconv.Atoi(authorIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author_id"})
			return
		}
		authorID = parsedAuthorID
	}

	if minRatingStr := c.Query("min_rating"); minRatingStr != "" {
		parsedMinRating, err := strconv.ParseFloat(minRatingStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid min_rating"})
			return
		}
		minRating = parsedMinRating
	}

	items, err := h.service.GetExplorer(authorID, minRating)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "author not found" {
			status = http.StatusNotFound
		}
		if err.Error() == "min_rating must be between 0 and 5" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

// GET /books/:id
func (h *BookHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	book, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
func (h *BookHandler) Create(c *gin.Context) {
	var req entities.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.Create(req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "author not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// PUT /books/:id
func (h *BookHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req entities.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.Update(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "book not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
func (h *BookHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "book not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}
