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

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgconn"
)

type JSON map[string]any
type ctxKey string

const (
	ctxUserIdKey    ctxKey = "userId"
	ctxUserEmailKey ctxKey = "userEmail"
	ctxUserRoleKey  ctxKey = "userRole"
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
		// get authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "no authorization header"})
			return
		}
		parts := strings.Fields(authHeader)
		if !strings.EqualFold(parts[0], "Bearer") {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "authorization must start with 'Bearer'"})
			return
		}
		if len(parts) != 2 {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "authorization must be formatted as 'Bearer <token>'"})
			return
		}
		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(model.JwtSecret), nil
		})
		// log.Printf("%v\n", token)
		// log.Printf("%v\n", err)
		if err != nil || !token.Valid {
			WriteJSON(w, http.StatusUnauthorized, "invalid token")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			WriteJSON(w, http.StatusUnauthorized, "invalid token claims")
			return
		}
		userId, ok := claims["userId"].(string)
		if !ok || userId == "" {
			WriteJSON(w, http.StatusUnauthorized, "missing userId in token")
			return
		}
		userEmail, ok := claims["userEmail"].(string)
		if !ok || userEmail == "" {
			WriteJSON(w, http.StatusUnauthorized, "missing userEmail in token")
			return
		}
		userRole, ok := claims["userRole"].(string)
		if !ok || userRole == "" {
			WriteJSON(w, http.StatusUnauthorized, "missing userRole in token")
			return
		}
		// inject into context
		ctx := context.WithValue(r.Context(), ctxUserIdKey, userId)
		ctx = context.WithValue(ctx, ctxUserEmailKey, userEmail)
		ctx = context.WithValue(ctx, ctxUserRoleKey, userRole)
		// TODO: query database to check if user is active?
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Authorization middleware
func (s *Server) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value(ctxUserRoleKey).(string)
		if userRole != string(model.RoleAdmin) {
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
	var userDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}

	userExisted, err := s.db.SelectUserByEmail(userDTO.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "email not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	log.Println(userExisted)
	if !userExisted.ValidatePassword(userDTO.Password) {
		WriteJSON(w, http.StatusUnauthorized,
			JSON{"error": "password incorrect"},
		)
		return
	}

	token, err := userExisted.GenerateJWT()
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
	userId := r.Context().Value(ctxUserIdKey).(string)
	userEmail := r.Context().Value(ctxUserEmailKey).(string)
	existedUser, err := s.db.SelectUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "userid in token not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if existedUser.Email != userEmail {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "token is corrupted - id and email mismatch"})
		return
	}
	if !existedUser.IsActive {
		WriteJSON(w, http.StatusNotFound, JSON{"error": "user is not active"})
		return
	}
	WriteJSON(w, http.StatusOK, existedUser.ToUserDTO())
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	// check if incoming email is valid
	if err := user.IsValidEmail(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	// no need to check if type assertions are successful because we did that in
	// authenticate middleware already
	userId := r.Context().Value(ctxUserIdKey).(string)
	userEmail := r.Context().Value(ctxUserEmailKey).(string)
	// query database to check if user exists and is_active before updating
	existedUser, err := s.db.SelectUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "userid in token not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if existedUser.Email != userEmail {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "token is corrupted - id and email mismatch"})
		return
	}
	if !existedUser.IsActive {
		WriteJSON(w, http.StatusNotFound, JSON{"error": "user is not active"})
		return
	}
	updatedUser, err := s.db.UpdateUserEmail(userId, &user)
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
	WriteJSON(w, http.StatusOK, updatedUser.ToUserDTO())
}

func (s *Server) StatusUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id       string `json:"id"`
		IsActive bool   `json:"is_active"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decode request body: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	userId := r.Context().Value(ctxUserIdKey).(string)
	userEmail := r.Context().Value(ctxUserEmailKey).(string)
	userRole := r.Context().Value(ctxUserRoleKey).(string)
	existedUser, err := s.db.SelectUserById(body.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "user to be set status not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	// admin can activate or deactivate any user
	// user can only deactivate itself
	if userRole != string(model.RoleAdmin) {
		// activate
		if body.IsActive {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "only admin can activate a user"})
			return
		}
		// deactivate
		if userId != existedUser.Id {
			WriteJSON(w, http.StatusForbidden, JSON{"error": "only user can deactivate itself"})
			return
		}
		// other edge case, when auth user is not admin, but the email in token
		// don't match
		if userEmail != existedUser.Email {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "token is corrupted - id and email mismatch"})
			return
		}
		// fine
	}
	err = s.db.UpdateUserStatus(body.Id, body.IsActive)
	if err != nil {
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
	if strings.TrimSpace(body.OldPassword) == "" {
		WriteJSON(w, http.StatusBadRequest, JSON{"error": "old password cannot be empty"})
		return
	}
	// tmpUser to User methods like HashPassword and ValidatePassword
	tmpUser := model.User{Password: body.NewPassword}
	if err := tmpUser.IsValidPassword(); err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusBadRequest, JSON{"error": err.Error()})
		return
	}
	userId := r.Context().Value(ctxUserIdKey).(string)
	userEmail := r.Context().Value(ctxUserEmailKey).(string)
	existedUser, err := s.db.SelectUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusUnauthorized, JSON{"error": "userid in token not found"})
			return
		}
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	if existedUser.Email != userEmail {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "token is corrupted - id and email mismatch"})
		return
	}
	if !existedUser.IsActive {
		WriteJSON(w, http.StatusNotFound, JSON{"error": "user is not active"})
		return
	}
	if !existedUser.ValidatePassword(body.OldPassword) {
		WriteJSON(w, http.StatusUnauthorized, JSON{"error": "old password is not correct"})
		return
	}
	err = s.db.UpdateUserPassword(userId, &tmpUser)
	if err != nil {
		log.Printf("Error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {}
