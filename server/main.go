package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port :9000")
	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../assets"))))
}
