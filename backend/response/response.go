package response

import "github.com/gin-gonic/gin"

func RespondSuccess(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data, "message": "success"})
}

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
