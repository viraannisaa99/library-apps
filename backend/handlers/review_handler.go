package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/services"
)

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

// GET /reviews
func (h *ReviewHandler) GetAll(c *gin.Context) {
	// Support filter by book: GET /reviews?book_id=1
	if bookIDStr := c.Query("book_id"); bookIDStr != "" {
		bookID, err := strconv.Atoi(bookIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book_id"})
			return
		}
		reviews, err := h.service.GetByBookID(bookID)
		if err != nil {
			status := http.StatusInternalServerError
			if err.Error() == "book not found" {
				status = http.StatusNotFound
			}
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": reviews})
		return
	}

	reviews, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// GET /reviews/:id
func (h *ReviewHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	review, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// POST /reviews
func (h *ReviewHandler) Create(c *gin.Context) {
	var req entities.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.Create(req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "book not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": review})
}

// PUT /reviews/:id
func (h *ReviewHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req entities.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.Update(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "review not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DELETE /reviews/:id
func (h *ReviewHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "review not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "review deleted"})
}
