package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"auth/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", NewServer.port),
		Handler:        NewServer.RegisterRoutes(),
		IdleTimeout:    time.Minute,
		MaxHeaderBytes: 1 << 10,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	return server
}
