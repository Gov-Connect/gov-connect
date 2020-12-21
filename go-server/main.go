package main

import (
	"fmt"
	"log"
	"net/http"

	"go-server/router"
)

func main() {
	//likely will switch to Gin or more direct net/http
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
