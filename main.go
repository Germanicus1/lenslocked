package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>To get in touch email me at <a href=\"mailto:peter.kerschbaumer.es\">peter@kerschbaumer.es</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/":
		homeHandler(w,r)
	case "/contact":
		contactHandler(w,r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}


func main() {
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", pathHandler)

	fmt.Println("Starting the server at :3000...")
	http.ListenAndServe(":3000", nil)
}


