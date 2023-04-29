package main

import (
	"chapter-history-api/handlers"
	"chapter-history-api/middleware"
	"chapter-history-api/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	db, err := NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	chapterVersionService := service.NewChapterVersionService(db)

	app.Use(middleware.RateLimiter(10, 2))
	app.Use(middleware.APIKeyAuth("123"))

	app.Get("/chapter_versions/:chapterID", handlers.GetChapterVersionsHandler(chapterVersionService))

	app.Listen(":9000")
}
