// Working with session cookies with the gorilla/sessions package
// Check the 'Application' devtool panel and check for cookies, you will find one called 'cookie-name'
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Create a new session store and add in the secret
var (
	key   = []byte("super-secret-key-12345")
	store = sessions.NewCookieStore(key)
)

func handleSecret(w http.ResponseWriter, r *http.Request) {
	// Get the session state for our cookie
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Otherwise, user is authenticated
	fmt.Fprintln(w, "You have special access to this thanks to being authenticated")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get session state for our cookie
	session, _ := store.Get(r, "cookie-name")

	// ... perform authentication (i.e. check username and password in db, etc)

	// Update session state that user is now authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	// Get session state for our cookie
	session, _ := store.Get(r, "cookie-name")

	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", handleSecret)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.ListenAndServe(":80", nil)
}
