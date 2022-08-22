package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}
