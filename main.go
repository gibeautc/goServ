package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting....")
	router:=NewRouter()
    log.Fatal(http.ListenAndServe(":8000", router))
}



