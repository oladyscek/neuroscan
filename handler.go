package main

import (
	"log"
	"net/http"
)

func main() {
	//createTable()
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//http.HandleFunc("/ws", handler)

	log.Println("server have started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}