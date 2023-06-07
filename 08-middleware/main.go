package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// A middleware in itself simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.

func main() {
	// Pass handlers through as arguments and wrap with middleware
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	// Using chain
	http.HandleFunc("/hello", Chain(Hello, Method("GET"), Logging()))

	http.ListenAndServe(":80", nil)
}

// Basic logging middleware
// Returns a func handler
// Calls a log statement before executing the handler
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// Other examples
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// Here we define a new type Middleware which makes it eventually easier to chain multiple middlewares together.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Example structure for creating a custom middleware function
func createMiddleware() Middleware {
	// Create a new middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandleFunc which s called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {
			// ...do middleware things

			// Call the next middlewaare/handler in chain
			next(w, r)
		}
		// Return the newly created handler
		return handler
	}
	// Return the newly created middleware
	return middleware
}

// Some custom middleware examples

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Log the request URL and the time it takes to run it
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			// next
			f(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Only accept provided VERB requests
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

// Chaining middleware
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
