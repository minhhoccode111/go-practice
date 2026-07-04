package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// Todo represents a todo item in the database.
type Todo struct {
	ID        int64
	Title     string
	Done      bool
	CreatedAt time.Time
}

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Migrate runs any pending database migrations.
	Migrate() error

	// CreateTodo saves a new todo and returns it.
	CreateTodo(title string) (Todo, error)

	// GetTodos returns all todos ordered by creation time.
	GetTodos() ([]Todo, error)

	// GetTodo returns a single todo by ID.
	GetTodo(id int64) (Todo, error)

	// UpdateTodo updates the title and returns the updated todo.
	UpdateTodo(id int64, title string) (Todo, error)

	// DeleteTodo removes a todo by ID.
	DeleteTodo(id int64) error

	// ToggleTodo toggles the done state and returns the updated todo.
	ToggleTodo(id int64) (Todo, error)
}

type service struct {
	db *sql.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// CreateTodo saves a new todo and returns it.
func (s *service) CreateTodo(title string) (Todo, error) {
	res, err := s.db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	if err != nil {
		return Todo{}, fmt.Errorf("create todo: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, fmt.Errorf("get last insert id: %w", err)
	}
	return Todo{ID: id, Title: title, Done: false, CreatedAt: time.Now()}, nil
}

// GetTodos returns all todos ordered by creation time.
func (s *service) GetTodos() ([]Todo, error) {
	rows, err := s.db.Query("SELECT id, title, done, created_at FROM todos ORDER BY created_at ASC")
	if err != nil {
		return nil, fmt.Errorf("query todos: %w", err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan todo: %w", err)
		}
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}
	return todos, nil
}

// GetTodo returns a single todo by ID.
func (s *service) GetTodo(id int64) (Todo, error) {
	var t Todo
	err := s.db.QueryRow("SELECT id, title, done, created_at FROM todos WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt)
	if err != nil {
		return Todo{}, fmt.Errorf("get todo: %w", err)
	}
	return t, nil
}

// UpdateTodo updates the title and returns the updated todo.
func (s *service) UpdateTodo(id int64, title string) (Todo, error) {
	_, err := s.db.Exec("UPDATE todos SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return Todo{}, fmt.Errorf("update todo: %w", err)
	}
	return s.GetTodo(id)
}

// DeleteTodo removes a todo by ID.
func (s *service) DeleteTodo(id int64) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}
	return nil
}

// ToggleTodo toggles the done state of a todo and returns the updated todo.
func (s *service) ToggleTodo(id int64) (Todo, error) {
	var t Todo
	err := s.db.QueryRow("SELECT id, title, done, created_at FROM todos WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt)
	if err != nil {
		return Todo{}, fmt.Errorf("get todo for toggle: %w", err)
	}

	_, err = s.db.Exec("UPDATE todos SET done = ? WHERE id = ?", !t.Done, id)
	if err != nil {
		return Todo{}, fmt.Errorf("toggle todo: %w", err)
	}
	t.Done = !t.Done
	return t, nil
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}

// Migrate runs any pending database migrations from embedded SQL files.
func (s *service) Migrate() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		filename TEXT PRIMARY KEY,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("create migrations table: %w", err)
	}

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	for _, f := range files {
		var applied int
		err := s.db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE filename = ?", f).
			Scan(&applied)
		if err != nil {
			return fmt.Errorf("check migration %s: %w", f, err)
		}
		if applied > 0 {
			continue
		}

		content, err := migrationsFS.ReadFile("migrations/" + f)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", f, err)
		}

		_, err = s.db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("apply migration %s: %w", f, err)
		}

		_, err = s.db.Exec("INSERT INTO schema_migrations (filename) VALUES (?)", f)
		if err != nil {
			return fmt.Errorf("record migration %s: %w", f, err)
		}

		log.Printf("Applied migration: %s", f)
	}

	return nil
}
