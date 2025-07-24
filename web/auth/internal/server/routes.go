package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"auth/internal/model"

	"github.com/gorilla/mux"
)

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

	// user-authenticated routes
	user := r.PathPrefix("/users").Subrouter()
	user.Use(s.authMiddleware)
	user.HandleFunc("/{id}", s.GetUserHandler).Methods("GET")
	user.HandleFunc("/{id}", s.UpdateUserHandler).Methods("PUT")
	user.HandleFunc("/{id}/password", s.PasswordUserHandler).Methods("PATCH")
	// user can deactivate their account
	// but only admin can activate an account
	user.HandleFunc("/{id}/status", s.StatusUserHandler).Methods("PATCH")

	// admin-authorized routes
	admin := r.PathPrefix("/users").Subrouter()
	admin.Use(s.authMiddleware)
	admin.Use(s.adminMiddleware)
	admin.HandleFunc("/all", s.GetUsersHandler).Methods("GET")
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
	WriteJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, s.db.Health())
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO model.User
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := userDTO.IsValidEmail(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := userDTO.IsValidPassword(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	userExisted, err := s.db.SelectUserByEmail(userDTO.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if userExisted != nil {
		WriteJSON(w, http.StatusConflict, map[string]string{"error": "email already existed"})
		return
	}
	err = s.db.InsertUser(&userDTO)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	token, err := userDTO.GenerateJWT()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusCreated, token)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO model.User
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	userExisted, err := s.db.SelectUserByEmail(userDTO.Email)
	if err != nil && err == sql.ErrNoRows {
		WriteJSON(w, http.StatusUnauthorized,
			map[string]string{"message": "email or password incorrect"},
		)
		return
	}
	if !userExisted.ValidatePassword(userDTO.Password) {
		WriteJSON(w, http.StatusUnauthorized,
			map[string]string{"message": "email or password incorrect"},
		)
		return
	}
	token, err := userDTO.GenerateJWT()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusCreated, token)
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request)     {}
func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request)      {}
func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) StatusUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) PasswordUserHandler(w http.ResponseWriter, r *http.Request) {}
func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request)   {}
