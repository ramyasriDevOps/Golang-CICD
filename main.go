package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler executed")
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server is starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting the server", err)
	}
}
