package main

import (
	"fmt"
	"log"

	"go-server/router"
)

func main() {

	r := router.Router()
	log.Fatal(r.Run(":8080"))
	fmt.Println("Starting server on the port 8080...")
}
