package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/response"
)

func ParseIDParam(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "invalid id")
		return 0, false
	}
	return id, true
}
