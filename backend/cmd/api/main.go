package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vira/go-crud/config"
	"github.com/vira/go-crud/handlers"
	"github.com/vira/go-crud/repositories"
	"github.com/vira/go-crud/server"
	"github.com/vira/go-crud/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file loaded: %v", err)
	}

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
	r := server.NewRouter(authorHandler, bookHandler, reviewHandler)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
