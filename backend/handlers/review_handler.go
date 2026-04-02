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

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

// GET /reviews
func (h *ReviewHandler) GetAll(c *gin.Context) {
	// Support filter by book: GET /reviews?book_id=1
	if bookIDStr := c.Query("book_id"); bookIDStr != "" {
		bookID, err := strconv.Atoi(bookIDStr)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "invalid book_id")
			return
		}
		reviews, err := h.service.GetByBookID(bookID)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, services.ErrBookNotFound) {
				status = http.StatusNotFound
			}
			response.RespondError(c, status, err.Error())
			return
		}
		response.RespondSuccess(c, http.StatusOK, reviews)
		return
	}

	reviews, err := h.service.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, reviews)
}

// GET /reviews/:id
func (h *ReviewHandler) GetByID(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	review, err := h.service.GetByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrReviewNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, review)
}

// POST /reviews
func (h *ReviewHandler) Create(c *gin.Context) {
	var req entities.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.service.Create(req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrBookNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusCreated, review)
}

// PUT /reviews/:id
func (h *ReviewHandler) Update(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	var req entities.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.service.Update(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrReviewNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, review)
}

// DELETE /reviews/:id
func (h *ReviewHandler) Delete(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrReviewNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, "review deleted")
}
