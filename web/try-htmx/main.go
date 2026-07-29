package main

import (
	"net/http"
	"os"

	"try-htmx/components"
	"try-htmx/layout"
	renderer "try-htmx/templrender"

	"github.com/a-h/templ"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	IsDone bool
	Title  string
}

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

	db.AutoMigrate(&Todo{})

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
			renderer.Render(r.Context(), w,
				templ.Join(layout.PageTitle("Home"), nil),
			)
			return
		}

		renderer.Render(
			r.Context(),
			w,
			layout.App(
				layout.AppData{PageTitle: "Home"},
				nil,
			),
		)
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		content := components.Hello("World")
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w,
				templ.Join(layout.PageTitle("Hello"), content),
			)
			return
		}

		renderer.Render(
			r.Context(),
			w,
			layout.App(
				layout.AppData{PageTitle: "Hello"},
				content,
			),
		)
	})

	r.Get("/counter", func(w http.ResponseWriter, r *http.Request) {
		content := components.Counter()
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w,
				templ.Join(layout.PageTitle("Counter"), content))
			return
		}

		renderer.Render(
			r.Context(),
			w,
			layout.App(
				layout.AppData{PageTitle: "Counter"},
				content,
			),
		)
	})

	if err := http.ListenAndServe(":"+c.Port, r); err != nil {
		panic(err)
	}
}
