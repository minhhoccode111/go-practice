package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userDb := os.Getenv("USERDB")
	passDb := os.Getenv("PASSDB")
	nameDb := os.Getenv("NAMEDB")

	if userDb == "" || passDb == "" || nameDb == "" {
		log.Fatal("Missing required env vars: USERDB, PASSDB, NAMEDB")
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=127.0.0.1 port=5432 sslmode=disable",
		userDb,
		passDb,
		nameDb,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	// hard-code id 2 here to test the query.
	album, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", album)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)
}

// addAlbum adds the specified album to the database, returning the album ID of
// the new entry
func addAlbum(a Album) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", a.Title, a.Artist, a.Price).Scan(&id)

	// check for an error from the attempt to INSERT
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	// if no error was returned, return the ID
	return id, nil
}

// albumByID queries for the albun with the specified ID.
func albumByID(id int64) (Album, error) {
	// an album to hold data from the returned row.
	var album Album

	// execute a SELECT statement to query for an album with the specified ID.
	// it returns an sql.Row to simplify the calling code (your code!),
	// QueryRow doesn't return an error. Instead, it arranges to return any query
	// error (such as sql.ErrNoRows) from Rows.Scan later
	row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)

	// use Row.Scan to copy column values into struct fields
	// check for an error from Scan
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		// the special error sql.ErrNoRows indicates that the query returned no
		// rows. Typically that error is worth replacing with more specific
		// text, such as "no such album" here
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumsById %d: no such album", id)
		}
		// other type of error
		return album, fmt.Errorf("albumsById %d: %v", id, err)
	}

	// if no error just return the album
	return album, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// an albums slice to hold data from returned rows.
	var albums []Album

	// execute a SELECT statement to query for albums with the specified artist
	// name.
	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// closing rows so that any resources it holds will be released when
	// function exists
	defer rows.Close()

	// loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var album Album
		// Scan takes a list of pointers to Go values, where the column values
		// will be written. Scan writes through the pointers to update the
		// struct fields
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		// then append the new album to the albums slice
		albums = append(albums, album)
	}

	// check for an error from the overall query, using rows.Err()
	// note the if the query itself fails, checking for an error here is the
	// only way to find out that the results are incomplete
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// return albums slice without error
	return albums, nil
}
