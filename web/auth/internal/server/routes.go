package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"auth/internal/model"

	"github.com/gorilla/mux"
)

type JSON map[string]any

func WriteText(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(status)
	w.Write([]byte(data))
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	// public
	r.HandleFunc("/", s.HelloWorldHandler)
	r.HandleFunc("/health", s.healthHandler)
	r.HandleFunc("/auth/register", s.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", s.LoginHandler).Methods("POST")
	// WARN: must define '/all' before '/{id}'
	r.HandleFunc("/users/all", s.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", s.GetUserHandler).Methods("GET")

	// user-authenticated routes
	user := r.PathPrefix("/users").Subrouter()
	user.Use(s.authMiddleware)
	user.HandleFunc("/{id}", s.UpdateUserHandler).Methods("PATCH")
	user.HandleFunc("/{id}/password", s.PasswordUserHandler).Methods("PATCH")
	// user can deactivate their account
	// but only admin can activate an account
	user.HandleFunc("/{id}/status", s.StatusUserHandler).Methods("PATCH")

	// admin-authorized routes
	admin := r.PathPrefix("/users").Subrouter()
	admin.Use(s.authMiddleware)
	admin.Use(s.adminMiddleware)
	admin.HandleFunc("/{id}", s.DeleteUserHandler).Methods("DELETE")

	return r
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Authentication middleware
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check jwt, set user in context
		next.ServeHTTP(w, r)
	})
}

// Authorization middleware
func (s *Server) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get user from context, check role == admin
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, JSON{"message": "Hello, World!"})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, s.db.Health())
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO model.User
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	if err := userDTO.IsValidEmail(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	if err := userDTO.IsValidPassword(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	userExisted, err := s.db.SelectUserByEmail(userDTO.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if userExisted != nil {
		WriteJSON(w, http.StatusConflict, JSON{"error": "email already existed"})
		return
	}
	err = s.db.InsertUser(&userDTO)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	token, err := userDTO.GenerateJWT()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusCreated, token)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO model.User
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}

	userExisted, err := s.db.SelectUserByEmail(userDTO.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "email notfound"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	fmt.Println(userDTO)
	if !userExisted.ValidatePassword(userDTO.Password) {
		WriteJSON(w, http.StatusUnauthorized,
			JSON{"error": "password incorrect"},
		)
		return
	}

	token, err := userDTO.GenerateJWT()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, token)
}

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	userid := paths[len(paths)-1]
	user, err := s.db.SelectUserById(userid)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, JSON{
				"error": "user not found",
			})
		} else {
			log.Printf("Error: %v", err)
		}
		return
	}
	WriteJSON(w, http.StatusOK, user.ToUserDTO())
}

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	perPageStr := r.URL.Query().Get("perPage")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	allStr := r.URL.Query().Get("all")
	filter := r.URL.Query().Get("q")
	var err error

	limit, err := strconv.Atoi(perPageStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}
	isGetAll := allStr == "true"
	offset := (pageNumber - 1) * limit
	users, err := s.db.SelectUsers(limit, offset, filter, isGetAll)
	if err != nil {
		log.Printf("error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	countUsers, err := s.db.CountUsers(filter, isGetAll)
	if err != nil {
		log.Printf("error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	divideAndRoundUp := func(a, b int) int {
		return (a + b - 1) / b // e.g. 10 / 3 = (10 + 3 - 1) / 3 = 4
	}

	WriteJSON(w, http.StatusOK, JSON{
		"users":      users,
		"totalPage":  divideAndRoundUp(countUsers, limit),
		"perPage":    limit,
		"pageNumber": pageNumber,
	})
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) StatusUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) PasswordUserHandler(w http.ResponseWriter, r *http.Request) {}
func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request)   {}
