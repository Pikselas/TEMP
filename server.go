package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<h1>Hello World</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
