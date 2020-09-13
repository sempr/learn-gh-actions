package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	type BlogData struct {
		Title   string
		Content string
	}
	resp, _ := json.Marshal([]*BlogData{
		{"First Blog", "bla bla bla..."},
		{"Second Blog", "Hello, world! "},
	})

	f, err := os.Create("./api")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	f.Write(resp)
}
