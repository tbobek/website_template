package main

import (
	"log"
	"net/http"

	"github.com/flosch/pongo2/v4"
)

// Pre-compiling the templates at application startup using the
// little Must()-helper function (Must() will panic if FromFile()
// or FromString() will return with an error - that's it).
// It's faster to pre-compile it anywhere at startup and only
// execute the template later.
var tplExample = pongo2.Must(pongo2.FromFile("./example.html"))

func examplePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received")
	// Execute the template per HTTP request
	log.Printf("Query: %v", r.URL.Query())
	err := tplExample.ExecuteWriter(pongo2.Context{"query": r.URL.Query()["query"], "name": "test"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", examplePage)
	log.Println("starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
