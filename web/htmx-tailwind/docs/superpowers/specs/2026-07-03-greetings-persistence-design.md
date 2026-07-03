# Greetings Persistence Design

## Goal
- `GET /web` loads existing greetings from SQLite and displays them (instead of empty container)
- `POST /hello` saves new greetings to the DB and appends them to the page

## Changes

### 1. Database (`internal/database/database.go`)
- Add `Greeting` struct: `ID int64`, `Name string`, `CreatedAt time.Time`
- Add `InsertGreeting(name string) (Greeting, error)` — INSERT into greetings table
- Add `GetGreetings() ([]Greeting, error)` — SELECT all ordered by created_at

### 2. Templates (`cmd/web/hello.templ`)
- Change `HelloForm(greetings []Greeting)` — accepts DB greetings, renders each as a `HelloPost` card inside `#hello-container`
- `HelloPost(name string)` — unchanged

### 3. Handler (`cmd/web/hello.go`)
- `HelloWebHandler` (POST): saves `name` to DB, then renders `HelloPost(name)` (returns HTML fragment for htmx to append)
- New `HelloWebGetHandler` (GET): loads `[]Greeting` from DB, renders `HelloForm(greetings)`

### 4. Routes (`internal/server/routes.go`)
- `GET /web` → `HelloWebGetHandler` (instead of direct template render)
- `POST /hello` → `HelloWebHandler` (unchanged route, but handler now saves to DB)

### 5. Cleanup
- Remove old `test.db`, regenerate templ, rebuild
