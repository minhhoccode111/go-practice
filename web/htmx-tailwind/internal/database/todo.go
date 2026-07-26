package database

import "time"

// Todo represents a todo item in the database.
type Todo struct {
	ID        int64
	Title     string
	Done      bool
	CreatedAt time.Time
}
