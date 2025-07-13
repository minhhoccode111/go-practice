package server

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"time"

	"github.com/gorilla/mux"

	"github.com/coder/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	r.HandleFunc("/", s.HelloWorldHandler)

	r.HandleFunc("/auth/register", s.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", s.LoginHandler).Methods("POST")

	r.HandleFunc("/users/all", s.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", s.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", s.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}/status", s.StatusUserHandler).Methods("PATCH")
	r.HandleFunc("/users/{id}/password", s.PasswordUserHandler).Methods("PATCH")
	r.HandleFunc("/users/{id}", s.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/health", s.healthHandler)

	r.HandleFunc("/websocket", s.websocketHandler)
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

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

// TODO:
func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request)     {}
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request)        {}
func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request)     {}
func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request)      {}
func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) StatusUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) PasswordUserHandler(w http.ResponseWriter, r *http.Request) {}
func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request)   {}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
