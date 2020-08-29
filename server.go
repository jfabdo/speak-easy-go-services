package main

import (
	"fmt"
	"log"
	"net/http"

	"./pages"
)

func main() {
	http.HandleFunc("/", pages.HandleSe)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
