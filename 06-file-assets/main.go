package main

import "net/http"

func main() {
	// Start a file sever at the ./assets/ directory
	fs := http.FileServer(http.Dir("assets/"))

	// define handler to use file server for given route, strip route to sync root paths
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start server to view files at the /static route
	http.ListenAndServe(":80", nil)
}
