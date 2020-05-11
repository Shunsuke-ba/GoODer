package main

import (
	"log"
	"net/http"
	"GoODer/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.GetTweetHandler())

	log.Println("Server Running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe is failed %v:", err)
	}
}