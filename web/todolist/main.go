package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

var todos = []Todo{
	{Id: 0, Description: "Init project todolist", IsDone: true},
	{Id: 1, Description: "Finish project todolist", IsDone: false},
}

var countId = len(todos)

func main() {
	port := ":8000"
	fmt.Println("Server is listening on port:", port)
	s := &http.Server{
		Addr:           port,
		Handler:        nil,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 10, // 1024
	}
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/todos", handleTodos)     // get, post
	http.HandleFunc("/todos/{id}", handleTodo) // get, put, delete
	log.Fatal(s.ListenAndServe())
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Hello, World!",
	})
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetTodos(w, r)
	case "POST":
		handlePostTodo(w, r)
	default:
		handleNotAllowed(w)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetTodo(w, r)
	case "PUT":
		handlePutTodo(w, r)
	case "DELETE":
		handleDeleteTodo(w, r)
	default:
		handleNotAllowed(w)
	}
}

func handleNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	perPageStr := r.URL.Query().Get("perPage")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	var err error

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	start := min((pageNumber-1)*perPage, len(todos))
	end := min(start+perPage, len(todos))
	divideAndRoundUp := func(a, b int) int {
		return (a + b - 1) / b // e.g. 10 / 3 = (10 + 3 - 1) / 3 = 4
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"totalPage":  divideAndRoundUp(len(todos), perPage),
		"perPage":    perPage,
		"pageNumber": pageNumber,
		"todos":      todos[start:end],
	})
}

func handlePostTodo(w http.ResponseWriter, r *http.Request) {
	var todoDTO Todo
	err := json.NewDecoder(r.Body).Decode(&todoDTO)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if todoDTO.Description == "" {
		http.Error(w, "Description cannot be empty", http.StatusBadRequest)
		return
	}

	newTodo := Todo{
		Id:          countId,
		Description: todoDTO.Description,
		IsDone:      false,
	}

	todos = append(todos, newTodo)
	countId++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	idStr := paths[len(paths)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Please provide valid Todo Id", http.StatusBadRequest)
		return
	}
	for _, v := range todos {
		if v.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func handlePutTodo(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	idStr := paths[len(paths)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Please provide valid Todo Id", http.StatusBadRequest)
		return
	}

	var todoDTO Todo
	json.NewDecoder(r.Body).Decode(&todoDTO)

	if todoDTO.Description == "" {
		http.Error(w, "Description cannot be empty", http.StatusBadRequest)
		return
	}

	for i, v := range todos {
		if v.Id == id {
			newTodo := Todo{
				Id:          v.Id, // can't update Todo.Id
				Description: todoDTO.Description,
				IsDone:      todoDTO.IsDone,
			}
			todos[i] = newTodo
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newTodo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	idStr := paths[len(paths)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Please provide valid Todo Id", http.StatusBadRequest)
		return
	}

	for i, v := range todos {
		if v.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}
