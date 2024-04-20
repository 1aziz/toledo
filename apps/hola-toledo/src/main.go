package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("service started..\n")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/halo", getHalo)
	http.ListenAndServe(":8080", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hola Toledo!!\n")
}

func getHalo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hola request\n")
	io.WriteString(w, "Hola, HTTP!\n")
}
