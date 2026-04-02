package server

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/handlers"
)

func NewRouter(
	authorHandler *handlers.AuthorHandler,
	bookHandler *handlers.BookHandler,
	reviewHandler *handlers.ReviewHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: getAllowedOrigins(),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	{
		authors := api.Group("/authors")
		{
			authors.GET("", authorHandler.GetAll)
			authors.GET("/:id", authorHandler.GetByID)
			authors.POST("", authorHandler.Create)
			authors.PUT("/:id", authorHandler.Update)
			authors.DELETE("/:id", authorHandler.Delete)
		}

		books := api.Group("/books")
		{
			books.GET("", bookHandler.GetAll)
			books.GET("/explorer", bookHandler.GetExplorer)
			books.GET("/:id", bookHandler.GetByID)
			books.POST("", bookHandler.Create)
			books.PUT("/:id", bookHandler.Update)
			books.DELETE("/:id", bookHandler.Delete)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("", reviewHandler.GetAll)
			reviews.GET("/:id", reviewHandler.GetByID)
			reviews.POST("", reviewHandler.Create)
			reviews.PUT("/:id", reviewHandler.Update)
			reviews.DELETE("/:id", reviewHandler.Delete)
		}
	}

	return r
}

func getAllowedOrigins() []string {
	defaultOrigins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
	}

	raw := strings.TrimSpace(os.Getenv("CORS_ORIGIN"))
	if raw == "" {
		return defaultOrigins
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, p := range parts {
		origin := strings.TrimSpace(p)
		if origin != "" {
			origins = append(origins, origin)
		}
	}

	if len(origins) == 0 {
		return defaultOrigins
	}

	return origins
}
