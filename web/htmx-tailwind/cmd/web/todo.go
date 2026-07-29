package web

import (
	"bytes"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"htmx-tailwind/internal/database"
)

func TodoGetHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		t, err := db.GetTodo(id)
		if err != nil {
			log.Printf("Error getting todo: %v", err)
			return c.Status(fiber.StatusNotFound).SendString("Not found")
		}

		component := TodoItem(t)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func TodoEditGetHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		t, err := db.GetTodo(id)
		if err != nil {
			log.Printf("Error getting todo: %v", err)
			return c.Status(fiber.StatusNotFound).SendString("Not found")
		}

		component := TodoEditForm(t)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func TodoUpdateHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		if err := c.BodyParser(c); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		title := c.FormValue("title")

		t, err := db.UpdateTodo(id, title)
		if err != nil {
			log.Printf("Error updating todo: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Update error")
		}

		component := TodoItem(t)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func ToggleTodoHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		t, err := db.ToggleTodo(id)
		if err != nil {
			log.Printf("Error toggling todo: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Toggle error")
		}

		component := TodoItem(t)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func TodoCreateHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(c); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		title := c.FormValue("title")

		t, err := db.CreateTodo(title)
		if err != nil {
			log.Printf("Error saving todo: %v", err)
		}

		component := TodoItem(t)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}

func TodoPageHandler(db database.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todos, err := db.GetTodos()
		if err != nil {
			log.Printf("Error loading todos: %v", err)
			todos = nil
		}

		component := TodoPage(todos)
		buf := new(bytes.Buffer)
		err = component.Render(c.Context(), buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Render error")
		}

		c.Type("html")
		return c.Status(fiber.StatusOK).SendString(buf.String())
	}
}
