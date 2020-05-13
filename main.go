package main

import (
	"GoODer/pkg/cron"
	"GoODer/pkg/handler"
	"log"
	"net/http"

	"github.com/carlescere/scheduler"
)

func main() {
	http.HandleFunc("/", handler.CreateFavoriteHandler())

	scheduler.Every(2).Hours().Run(cron.CreateFavoriteCron) // 2時間ごと
	log.Println("Server Running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe is failed %v:", err)
	}
}
