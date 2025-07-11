package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strconv"
	_ "strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

const (
	appPort        = ":8000"
	dbHost         = "localhost"
	dbPort         = 5432
	dbUser         = "mhc"
	dbPassword     = "Bruh0!0!"
	dbName         = "authz"
	dbSslMode      = "disable"
	readTimeout    = 1 * time.Second
	writeTimeout   = 1 * time.Second
	maxHeaderBytes = 1 << 10 // 1024
)

type IUser interface {
	GenerateJWT() string
}

// change to interface and add some methods
type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsActive  *bool     `json:"is_active"` // user pointer to differentiate between 'no provided' and 'explicitly false'
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string
}

var db *sql.DB

func main() {
	db = New()
	r := mux.NewRouter()

	// start server
	fmt.Println("Server is listening on port:", appPort)
	s := &http.Server{
		Handler:        r, // for routing
		Addr:           appPort,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes, // 1024
	}

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/auth/register", register).Methods("POST")
	r.HandleFunc("/auth/login", login).Methods("POST")

	r.HandleFunc("/users/all", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}/status", statusUser).Methods("PATCH")
	r.HandleFunc("/users/{id}/password", passwordUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(s.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, World!",
	})
}

func register(w http.ResponseWriter, r *http.Request)     {}
func login(w http.ResponseWriter, r *http.Request)        {}
func getUser(w http.ResponseWriter, r *http.Request)      {}
func updateUser(w http.ResponseWriter, r *http.Request)   {}
func statusUser(w http.ResponseWriter, r *http.Request)   {}
func passwordUser(w http.ResponseWriter, r *http.Request) {}
func deleteUser(w http.ResponseWriter, r *http.Request)   {}
