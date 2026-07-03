package server

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"htmx-tailwind/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	db := database.New()

	if err := db.Migrate(); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "htmx-tailwind",
			AppName:      "htmx-tailwind",
		}),

		db: db,
	}

	return server
}
