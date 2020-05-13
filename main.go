package main

import (
	"GoODer/config"
	"GoODer/pkg/cron"
	"GoODer/pkg/handler"
	"log"
	"net/http"

	"github.com/carlescere/scheduler"
	_ "github.com/carlescere/scheduler"
)

func main() {
	// バッチ処理用
	scheduler.Every(config.Config.FavoSleep).Hours().Run(cron.CreateFavoriteCron) // 3時間ごと

	// 手動で動かす用
	http.HandleFunc("/", handler.CreateFavoriteHandler())
	log.Println("Server Running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe is failed %v:", err)
	}

}
