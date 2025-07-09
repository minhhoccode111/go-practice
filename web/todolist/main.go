package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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

const (
	appPort        = ":8000"
	dbHost         = "localhost"
	dbPort         = 5432
	dbUser         = "mhc"
	dbPassword     = "Bruh0!0!"
	dbName         = "todolist"
	dbSslMode      = "disable"
	readTimeout    = 1 * time.Second
	writeTimeout   = 1 * time.Second
	maxHeaderBytes = 1 << 10 // 1024
)

var db *sql.DB

func main() {
	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

	var err error

	// create connection with the postgres db
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ping the db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// migrate or create tables
	_, err = db.Exec(`create table if not exists todos (
		id			serial	primary key,
		is_done		boolean default false,
		description	text	not		null
		);`)

	if err != nil {
		log.Fatal(err)
	}

	// generate fake data if doesn't have any
	var count int
	err = db.QueryRow("select count(*) from todos").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		_, err = db.Exec(
			`insert into todos (description, is_done) values ($1, $2), ($3, $4)`,
			"generate data test", true,
			"write unit tests", false,
		)
		if err != nil {
			log.Printf("Insert dummy data failed: %v", err)
		}
	}

	// start server
	fmt.Println("Server is listening on port:", appPort)
	s := &http.Server{
		Addr:           appPort,
		Handler:        nil,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes, // 1024
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

	limit, err := strconv.Atoi(perPageStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	offset := (pageNumber - 1) * limit

	rows, err := db.Query(`
		select id, description, is_done from todos
		order by id
		limit $1
		offset $2`,
		limit,
		offset,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Description, &todo.IsDone)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	var count int
	err = db.QueryRow(`select count(*) from todos`).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	divideAndRoundUp := func(a, b int) int {
		return (a + b - 1) / b // e.g. 10 / 3 = (10 + 3 - 1) / 3 = 4
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"totalPage":  divideAndRoundUp(count, limit),
		"perPage":    limit,
		"pageNumber": pageNumber,
		"todos":      todos,
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

	var newTodo Todo

	err = db.QueryRow(`
		insert into todos (description) values ($1)
		returning id, description, is_done
		`,
		todoDTO.Description,
	).Scan(&newTodo.Id, &newTodo.Description, &newTodo.IsDone)

	if err != nil {
		log.Fatal(err)
	}

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

	var todo Todo
	err = db.QueryRow(`select id, description, is_done from todos where id = $1`, id).Scan(&todo.Id, &todo.Description, &todo.IsDone)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}

		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
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

	var updatedTodo Todo
	err = db.QueryRow(`
		update todos
		set description = $1,
		is_done = $2
		where id = $3
		returning id, description, is_done
		`,
		todoDTO.Description,
		todoDTO.IsDone,
		id,
	).Scan(
		&updatedTodo.Id,
		&updatedTodo.Description,
		&updatedTodo.IsDone,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	idStr := paths[len(paths)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Please provide valid Todo Id", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`delete from todos where id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
