package main

import (
	"GoODer/pkg/cron"
	"GoODer/pkg/handler"
	"log"
	"net/http"

	"github.com/carlescere/scheduler"
)

func main() {
	// 手動で動かす用
	http.HandleFunc("/", handler.CreateFavoriteHandler())

	// バッチ処理用
	scheduler.Every(3).Minutes().Run(cron.CreateFavoriteCron) // 2時間ごと

	log.Println("Server Running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe is failed %v:", err)
	}

}
