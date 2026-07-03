package web

import (
	"bytes"
	"log"

	"github.com/gofiber/fiber/v2"

	"htmx-tailwind/internal/database"
)

func HelloWebHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(c); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		name := c.FormValue("name")

		_, err := db.InsertGreeting(name)
		if err != nil {
			log.Printf("Error saving greeting: %v", err)
		}

		component := HelloPost(name)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func HelloWebGetHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		greetings, err := db.GetGreetings()
		if err != nil {
			log.Printf("Error loading greetings: %v", err)
			greetings = nil
		}

		component := HelloForm(greetings)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		c.Type("html")
		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}
