package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	type BlogData struct {
		Title   string
		Content string
	}

	http.Handle("/", http.FileServer(http.Dir(".")))

	resp, _ := json.Marshal([]*BlogData{
		{"First Blog", "bla bla bla..."},
		{"Second Blog", "Hello, world! "},
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write(resp)
	})

	http.ListenAndServe(":8080", nil)
}
