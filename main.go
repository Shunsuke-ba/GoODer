package main

import (
	"GoODer/pkg/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.CreateFavoriteHandler())

	log.Println("Server Running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe is failed %v:", err)
	}
}
