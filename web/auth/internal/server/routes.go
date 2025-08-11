package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"auth/internal/model"
	"auth/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgconn"
)

type JSON map[string]any
type ctxKey string

const (
	ctxUserKey   ctxKey = "user"
	ctxUserIdKey ctxKey = "userId"
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
	r.HandleFunc("/healthz", s.healthHandler)
	r.HandleFunc("/auth/register", s.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", s.LoginHandler).Methods("POST")
	// WARN: must define '/all' before '/{id}'
	r.HandleFunc("/users/all", s.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", s.GetUserHandler).Methods("GET")

	// user-authenticated routes
	me := r.PathPrefix("/auth").Subrouter()
	me.Use(s.authMiddleware)
	me.HandleFunc("/me", s.GetMeHandler).Methods("GET")
	user := r.PathPrefix("/users").Subrouter()
	user.Use(s.authMiddleware)
	user.HandleFunc("/{id}", s.UpdateUserHandler).Methods("PATCH")
	user.HandleFunc("/{id}/password", s.PasswordUserHandler).Methods("PATCH")
	// WARN: user can deactivate their account but only admin can activate an account
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
		// get authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "no authorization header"})
			return
		}
		parts := strings.Fields(authHeader)
		if !strings.EqualFold(parts[0], "Bearer") {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "authorization header must start with 'Bearer'"})
			return
		}
		if len(parts) != 2 {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "authorization header must be formatted as 'Bearer <token>'"})
			return
		}
		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(utils.JwtSecret), nil
		})
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": err.Error()})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "invalid token"})
			return
		}
		userId, ok := claims[string(ctxUserIdKey)].(string)
		if !ok || userId == "" {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "missing userId in token"})
			return
		}
		user, err := s.db.SelectUserById(userId)
		if err != nil {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "cannot authorize user in jwt"})
			return
		}
		if !user.IsActive {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "user in jwt is inactive"})
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserIdKey, userId)
		ctx = context.WithValue(ctx, ctxUserKey, *user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Authorization middleware
func (s *Server) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(model.User)
		if user.Role != model.RoleAdmin {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "user is not admin"})
			return
		}
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
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	email, err := utils.IsValidEmail(body.Email)
	if err != nil {
		log.Printf("Input Email Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	password, err := utils.IsValidPassword(body.Password)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	userExisted, err := s.db.SelectUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if userExisted != nil {
		WriteJSON(w, http.StatusConflict, JSON{"error": "email already existed"})
		return
	}
	user := model.User{
		Email:    email,
		Password: password,
		IsActive: true,
		Role:     model.RoleUser,
	}
	err = s.db.InsertUser(&user)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	userDTO := utils.UserToUserDTO(&user)
	token, err := utils.GenerateJWT(&userDTO)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusCreated, JSON{"user": userDTO, "token": token})
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}

	userExisted, err := s.db.SelectUserByEmail(body.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "email not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if !utils.ValidatePassword(userExisted.Password, body.Password) {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "password incorrect"})
		return
	}
	userDTO := utils.UserToUserDTO(userExisted)
	token, err := utils.GenerateJWT(&userDTO)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, JSON{"user": userDTO, "token": token})
}

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	userId := paths[len(paths)-1] // path/user/{userId}
	existedUser, err := s.db.SelectUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, JSON{"error": "user not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, utils.UserToUserDTO(existedUser))
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
	// TODO: implement concurrency with goroutines and channels instead of running one by one like this
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

func (s *Server) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxUserKey).(model.User)
	userDTO := utils.UserToUserDTO(&user)
	token, err := utils.GenerateJWT(&userDTO)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, JSON{"token": token, "user": userDTO})
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email string `json:"email"` // NOTE: explicitly state what we will update
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	email, err := utils.IsValidEmail(body.Email)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	paths := strings.Split(r.URL.Path, "/")
	userIdPath := paths[len(paths)-1] // path/user/{userId}
	userIdToken := r.Context().Value(ctxUserIdKey).(string)
	if userIdPath != userIdToken {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "userIdToken and userIdPath mismatch"})
		return
	}
	updatedUserDTO, err := s.db.UpdateUser(userIdPath, email)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				WriteJSON(w, http.StatusConflict, JSON{"error": "email already existed"})
				return
			}
		}
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": "failed to update user profile"})
		log.Printf("Error: %v", err)
		return
	}
	WriteJSON(w, http.StatusOK, updatedUserDTO)
}

func (s *Server) StatusUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		// NOTE: pointer to differentiate between explicit-false and not-provided
		IsActive *bool `json:"is_active"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	if body.IsActive == nil {
		WriteJSON(w, http.StatusBadRequest, JSON{"error": "is_active is required"})
		return
	}
	paths := strings.Split(r.URL.Path, "/")
	userIdPath := paths[len(paths)-2] // path/users/{userId}/status
	userIdToken := r.Context().Value(ctxUserIdKey).(string)
	userInToken := r.Context().Value(ctxUserKey).(model.User)
	// admin can activate or deactivate any user, user can only deactivate itself
	if userInToken.Role != model.RoleAdmin {
		// activate
		if *body.IsActive {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "only admin can activate a user"})
			return
		}
		// deactivate
		if userIdToken != userIdPath {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "you must be admin to deactivate other users than yourself"})
			return
		}
		// fine to continue
	}
	err := s.db.UpdateUserStatus(userIdPath, *body.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "user to be set status not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) PasswordUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	newPassword, err := utils.IsValidPassword(body.NewPassword)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	paths := strings.Split(r.URL.Path, "/")
	userIdPath := paths[len(paths)-2] // path/users/{userId}/status
	userIdToken := r.Context().Value(ctxUserIdKey).(string)
	if userIdPath != userIdToken {
		WriteJSON(w, http.StatusForbidden, JSON{"error": "cannot change another user's password"})
		return
	}
	userInToken := r.Context().Value(ctxUserKey).(model.User)
	if !utils.ValidatePassword(userInToken.Password, body.OldPassword) {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "old password is not correct"})
		return
	}
	err = s.db.UpdateUserPassword(userIdPath, newPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "user to be updated password not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	userIdPath := paths[len(paths)-1] // path/users/{userId}
	userIdToken := r.Context().Value(ctxUserIdKey).(string)
	if userIdPath == userIdToken {
		WriteJSON(w, http.StatusForbidden, JSON{"error": "admin cannot self-delete"})
		return
	}
	err := s.db.DeleteUserById(userIdPath)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, JSON{"error": "user to be deleted not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}
