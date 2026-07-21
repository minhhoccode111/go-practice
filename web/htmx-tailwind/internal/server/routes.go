package server

import (
	"net/http"

	"htmx-tailwind/cmd/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

	s.App.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Files),
		PathPrefix: "assets",
		Browse:     false,
	}))

	s.App.Get("/todos", web.TodoPageHandler(s.db))

	s.App.Post("/todos", web.TodoCreateHandler(s.db))

	s.App.Get("/todos/:id", web.TodoGetHandler(s.db))

	s.App.Get("/todos/:id/edit", web.TodoEditGetHandler(s.db))

	s.App.Put("/todos/:id", web.TodoUpdateHandler(s.db))

	s.App.Patch("/todos/:id/toggle", web.ToggleTodoHandler(s.db))

	s.App.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}
		if err := s.db.DeleteTodo(int64(id)); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Delete error")
		}
		return c.SendString("")
	})
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
