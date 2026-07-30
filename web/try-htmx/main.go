package main

import (
	"net/http"
	"os"
	"strconv"

	"try-htmx/components"
	"try-htmx/entity"
	"try-htmx/layout"
	renderer "try-htmx/templrender"

	"github.com/a-h/templ"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port string
	DB   string
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db := os.Getenv("DB")
	if db == "" {
		db = "dev.db"
	}
	return Config{
		Port: port,
		DB:   db,
	}
}

func main() {
	// config
	c := loadConfig()

	// database
	db, err := gorm.Open(sqlite.Open(c.DB), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&entity.Todo{})

	// router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// serve files
	fileServer := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fileServer))

	// handlers
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w, templ.Join(layout.PageTitle("Home"), nil))
			return
		}

		renderer.Render(r.Context(), w, layout.App(layout.AppData{PageTitle: "Home"}, nil))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		content := components.Hello("World")
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w, templ.Join(layout.PageTitle("Hello"), content))
			return
		}

		renderer.Render(r.Context(), w, layout.App(layout.AppData{PageTitle: "Hello"}, content))
	})

	// counter -----------------------------------------------------------------

	counter := 0

	r.Get("/counter", func(w http.ResponseWriter, r *http.Request) {
		content := components.Counter(counter)
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w, templ.Join(layout.PageTitle("Counter"), content))
			return
		}

		renderer.Render(r.Context(), w, layout.App(layout.AppData{PageTitle: "Counter"}, content))
	})

	r.Put("/htmx/inc", func(w http.ResponseWriter, r *http.Request) {
		counter++
		w.Write([]byte(strconv.Itoa(counter)))
	})
	r.Put("/htmx/dec", func(w http.ResponseWriter, r *http.Request) {
		if counter > 0 {
			counter--
		}
		w.Write([]byte(strconv.Itoa(counter)))
	})
	r.Put("/htmx/reset", func(w http.ResponseWriter, r *http.Request) {
		counter = 0
		w.Write([]byte(strconv.Itoa(counter)))
	})

	// todos -------------------------------------------------------------------

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		content := layout.TodosLayout()
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w, templ.Join(layout.PageTitle("Todos"), content))
			return
		}

		renderer.Render(r.Context(), w, layout.App(layout.AppData{PageTitle: "Todos"}, content))
	})

	r.Post("/htmx/todos", func(w http.ResponseWriter, r *http.Request) {
	})

	r.Get("/htmx/todos", func(w http.ResponseWriter, r *http.Request) {
	})

	if err := http.ListenAndServe(":"+c.Port, r); err != nil {
		panic(err)
	}
}
