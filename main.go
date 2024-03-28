package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("initializing server")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/data", worldHandler)
	http.ListenAndServe(":3333", nil)
}

func worldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<p>hello world</p>")
}
