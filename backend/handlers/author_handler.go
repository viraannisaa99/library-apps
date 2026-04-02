package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/entities"
	"github.com/vira/go-crud/response"
	"github.com/vira/go-crud/services"
	"github.com/vira/go-crud/utils"
)

type AuthorHandler struct {
	service *services.AuthorService
}

func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{service}
}

// GET /authors
func (h *AuthorHandler) GetAll(c *gin.Context) {
	authors, err := h.service.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, authors)
}

// GET /authors/:id
func (h *AuthorHandler) GetByID(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	author, err := h.service.GetByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrAuthorNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, author)
}

// POST /authors
func (h *AuthorHandler) Create(c *gin.Context) {
	var req entities.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	author, err := h.service.Create(req)
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusCreated, author)
}

// PUT /authors/:id
func (h *AuthorHandler) Update(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	var req entities.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	author, err := h.service.Update(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrAuthorNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, author)
}

// DELETE /authors/:id
func (h *AuthorHandler) Delete(c *gin.Context) {
	id, ok := utils.ParseIDParam(c)
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrAuthorNotFound) {
			status = http.StatusNotFound
		}
		response.RespondError(c, status, err.Error())
		return
	}
	response.RespondSuccess(c, http.StatusOK, "author deleted")
}
