package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vira/go-crud/config"
	"github.com/vira/go-crud/handlers"
	"github.com/vira/go-crud/repositories"
	"github.com/vira/go-crud/services"
)

func main() {
	// 1. Connect DB
	db := config.NewDB()
	defer db.Close()

	// 2. Init Repositories
	authorRepo := repositories.NewAuthorRepository(db)
	bookRepo := repositories.NewBookRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)

	// 3. Init Services (inject repo dependencies)
	authorService := services.NewAuthorService(authorRepo)
	bookService := services.NewBookService(bookRepo, authorRepo)
	reviewService := services.NewReviewService(reviewRepo, bookRepo)

	// 4. Init Handlers
	authorHandler := handlers.NewAuthorHandler(authorService)
	bookHandler := handlers.NewBookHandler(bookService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// 5. Setup Router
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
		// Authors
		authors := api.Group("/authors")
		{
			authors.GET("", authorHandler.GetAll)
			authors.GET("/:id", authorHandler.GetByID)
			authors.POST("", authorHandler.Create)
			authors.PUT("/:id", authorHandler.Update)
			authors.DELETE("/:id", authorHandler.Delete)
		}

		// Books
		books := api.Group("/books")
		{
			books.GET("", bookHandler.GetAll) // GET /books?author_id=1 juga bisa
			books.GET("/explorer", bookHandler.GetExplorer)
			books.GET("/:id", bookHandler.GetByID)
			books.POST("", bookHandler.Create)
			books.PUT("/:id", bookHandler.Update)
			books.DELETE("/:id", bookHandler.Delete)
		}

		// Reviews
		reviews := api.Group("/reviews")
		{
			reviews.GET("", reviewHandler.GetAll) // GET /reviews?book_id=1 juga bisa
			reviews.GET("/:id", reviewHandler.GetByID)
			reviews.POST("", reviewHandler.Create)
			reviews.PUT("/:id", reviewHandler.Update)
			reviews.DELETE("/:id", reviewHandler.Delete)
		}
	}

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getAllowedOrigins() []string {
	// Default aman untuk local dev kalau env belum diset
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
