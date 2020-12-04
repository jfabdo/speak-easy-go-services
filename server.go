package main

import (
	"fmt"
	"log"
	"net/http"

	services "github.com/jfabdo/speak-easy-go-webserver/services/pages/handleSe"
)

func main() {
	http.HandleFunc("/", services.HandleSe)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
