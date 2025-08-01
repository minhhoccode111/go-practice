package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format("3:04:05PM")
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func main() {
	addr := "localhost:3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)
	// wrap entire mux with logger middleware
	// wrappedMux := NewLogger(mux)
	// wrap entire mux with logger and response header middleware
	wrappedMux := NewLogger(NewResponseHeader(mux, "Tai-Vi-Sao", "tai vi sao"))

	log.Printf("server is listening at %s", addr)
	// use wrappedMux instead of mux as root handler
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// ServeHttp() handles the request by passing it to the real handler and logging
// the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handler: handlerToWrap}
}

// ResponseHeader is a middleware handler that adds a header to the response
type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

// NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handlerToWrap http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}

// ServeHTTP handles the request by adding the response header
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// add the header
	w.Header().Add(rh.headerName, rh.headerValue)
	// call the wrapped handler
	rh.handler.ServeHTTP(w, r)
}

type EnsureAuth struct{ handler http.Handler }
type User struct{ username string }
type contextKey int

const authenticatedUserKey contextKey = 0

func GetAuthenticatedUser(r *http.Request) (*User, error) {
	// mock, do something with `r`
	return &User{"minhhoccode111"}, nil
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}

	// create a new request context containing the authenticated user
	ctxWithUser := context.WithValue(r.Context(), authenticatedUserKey, user)
	// create a new request using that new context
	rWithUser := r.WithContext(ctxWithUser)
	// call the real handler, passing the new request
	ea.handler.ServeHTTP(w, rWithUser)
}

func NewEnsureAuth(handlerToWrap http.Handler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}
