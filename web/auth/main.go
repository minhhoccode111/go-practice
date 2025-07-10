package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strconv"
	_ "strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

const (
	appPort        = ":8000"
	dbHost         = "localhost"
	dbPort         = 5432
	dbUser         = "mhc"
	dbPassword     = "Bruh0!0!"
	dbName         = "authz"
	dbSslMode      = "disable"
	readTimeout    = 1 * time.Second
	writeTimeout   = 1 * time.Second
	maxHeaderBytes = 1 << 10 // 1024
)

// change to interface and add some methods
type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string
}

func HashPassword(password string) string {
	// ignore error for simplicity
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

var db *sql.DB

func main() {
	connStr := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode,
	)

	var err error

	// create connection with the postgres db
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Create connection failed: ", err)
	}
	defer db.Close()

	// ping the db
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping database failed: ", err)
	}

	// create enum
	_, err = db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE user_role AS ENUM ('admin', 'user');
			END IF;
		END $$;
		`)
	if err != nil {
		log.Fatal("Create database enum type failed: ", err)
	}

	// create table TODO: add created_at and updated_at
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		id         UUID         PRIMARY KEY   DEFAULT gen_random_uuid(),
		role       user_role    NOT     NULL                           ,
		email      VARCHAR(255) NOT     NULL  UNIQUE                   ,
		password   VARCHAR(255) NOT     NULL                           ,
		is_active  BOOLEAN      DEFAULT FALSE                          ,
		last_name  VARCHAR(100) NOT     NULL                           ,
		first_name VARCHAR(100) NOT     NULL
		);
		`)
	if err != nil {
		log.Fatal("Create table users failed: ", err)
	}

	// create indexes on email, first_name, last_name
	_, err = db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email);
		CREATE INDEX IF NOT EXISTS idx_users_fullname ON users (first_name, last_name);
		`)
	if err != nil {
		// WARN: Create indexes failed: pq: column "email" does not exist
		// solved: drop table if exists users;
		log.Fatal("Create indexes failed: ", err)
	}

	// generate fake data if doesn't have any
	var count int
	err = db.QueryRow("select count(*) from users").Scan(&count)
	if err != nil {
		log.Fatal("Count rows in users failed: ", err)
	}
	// TODO: add created_at and updated_at
	if count == 0 {
		_, err = db.Exec(
			`
			insert into users
			(email, password, last_name, first_name, is_active, role)
			values
			($1, $2, $3, $4, $5, $6),
			($7, $8, $9, $10, $11, $12)
			`,
			"minhhoccode111@gmail.com", HashPassword("asdasd"), "Minh", "Dang", true, "admin",
			"asd@gmail.com", HashPassword("asdasd"), "Dummy", "Account", true, "user",
		)
		if err != nil {
			log.Printf("Insert dummy data failed:  %v", err)
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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", authHandler)   // get, post
	http.HandleFunc("/users", usersHandler) // get, put, delete
	log.Fatal(s.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

}
func authHandler(w http.ResponseWriter, r *http.Request) {

}
func usersHandler(w http.ResponseWriter, r *http.Request) {

}
