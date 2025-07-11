package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	// ignore error for simplicity
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func New() *sql.DB {
	var db *sql.DB
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

	// create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		id         UUID         PRIMARY KEY   DEFAULT gen_random_uuid(),
		role       user_role    NOT     NULL                           ,
		email      VARCHAR(255) NOT     NULL  UNIQUE                   ,
		password   VARCHAR(255) NOT     NULL                           ,
		is_active  BOOLEAN      DEFAULT FALSE                          ,
		last_name  VARCHAR(100) NOT     NULL                           ,
		first_name VARCHAR(100) NOT     NULL                           ,
		created_at TIMESTAMPTZ                DEFAULT now()            ,
		updated_at TIMESTAMPTZ                DEFAULT now()
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
	if count == 0 {
		_, err = db.Exec(
			`
			insert into users
			(email, password, last_name, first_name, is_active, role)
			values
			($1, $2, $3, $4, $5, $6),
			-- ($13, $14, $15, $16, $17, $18),
			($7, $8, $9, $10, $11, $12)
			`,
			"minhhoccode111@gmail.com", HashPassword("asdasd"), "Minh", "Dang", true, "admin",
			"asd@gmail.com", HashPassword("asdasd"), "Dummy", "Account", true, "user",
			// "asdasd@gmail.com", HashPassword("asdasd"), "Invalid", "Account", true, "invalid role", // test invalid role
		)
		if err != nil {
			log.Printf("Insert dummy data failed:  %v", err)
		}
	}
	return db
}
