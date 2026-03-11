package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Role string

const (
	admin Role = "admin"
	user  Role = "user"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:50;uniqueIndex;not null"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	Role     Role
	Bio      string `gorm:"type:text"`
	Image    string `gorm:"size:255"`

	Todos []Todo
}

type Todo struct {
	gorm.Model

	Title       string
	Description string
	IsDone      bool
	StartDate   time.Time
	DueDate     time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}
