# Go Web Authentication and Authorization

Concepts:

- Router
- Filter
- Context
- Interface
- Middleware
- Pagination
- Concurrency
- Soft-delete
- Authorization
- Authentication
- Use `PATCH` verb
- Database indexes
- Password hashing
- Database migration
- Input validation + sanitization
- Use pointer to differentiate between "no provided" and "explicitly false" in Go

Features:

APIs:

```rust
# public
POST   /auth/register
POST   /auth/login
GET    /users/all
    - filter email, status
    - pagination
    - TODO: only admin can get all users including inactive
GET    /users/{id}

# auth
PATCH  /users/{id}
    - user update their profile, like email, bio, etc.
PATCH  /users/{id}/password
    - user update their password, must provide old password
PATCH  /users/{id}/status
    - user deactivate her account (soft-delete)
    - only admin can activate an account

# authz
DELETE /users/{id}
    - admin hard-delete an account
PATCH  /users/{id}/role
    - admin make a user an admin
```

Database:

```txt
- Users:
    - id
    - role
idx - email     - unique
    - password
    - is_active
```

Todo:

- use UserDTO to transfer data most of the time
- use goroutines and channels to run queries concurrently
- use context in middlewares and database query
- refactor
- remove password from select user by id and email

```go
import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	perPageStr := r.URL.Query().Get("perPage")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	allStr := r.URL.Query().Get("all")
	filter := r.URL.Query().Get("q")
	var err error

	limit, err := strconv.Atoi(perPageStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}
	isGetAll := allStr == "true"
	offset := (pageNumber - 1) * limit

	// Channels to receive results from the goroutines
	usersCh := make(chan []User)
	countCh := make(chan int)
	errCh := make(chan error, 2) // Buffer for two possible errors

	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines

	// Goroutine for the SelectUsers query
	go func() {
		defer wg.Done()
		users, err := s.db.SelectUsers(limit, offset, filter, isGetAll)
		if err != nil {
			errCh <- err
			return
		}
		usersCh <- users
	}()

	// Goroutine for the CountUsers query
	go func() {
		defer wg.Done()
		countUsers, err := s.db.CountUsers(filter, isGetAll)
		if err != nil {
			errCh <- err
			return
		}
		countCh <- countUsers
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close channels after receiving the results
	close(usersCh)
	close(countCh)
	close(errCh)

	// Check for errors
	select {
	case err := <-errCh:
		log.Printf("error: %v", err)
		WriteJSON(w, http.StatusInternalServerError, JSON{"error": err.Error()})
		return
	default:
		// No error, continue processing
	}

	// Get the results from the channels
	users := <-usersCh
	countUsers := <-countCh

	// Calculate total pages
	divideAndRoundUp := func(a, b int) int {
		return (a + b - 1) / b // e.g. 10 / 3 = (10 + 3 - 1) / 3 = 4
	}

	// Respond with JSON data
	WriteJSON(w, http.StatusOK, JSON{
		"users":      users,
		"totalPage":  divideAndRoundUp(countUsers, limit),
		"perPage":    limit,
		"pageNumber": pageNumber,
	})
}
```
