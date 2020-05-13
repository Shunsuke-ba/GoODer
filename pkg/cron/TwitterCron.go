package cron

import (
	"GoODer/config"
	"GoODer/pkg/model"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func CreateFavoriteCron() {
	api, err := model.ConnectTwitterApi()
	if err != nil {
		log.Fatal("could not connect API")
	}
	log.Print("API connect is success")

	// 指定語句のツイートをGET(15件？)
	searchTweet, err := api.GetSearch(config.Config.FavoWord, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Print("search tweet is failed")
	}

	// 既にいいねしてるツイート取得（20件？）
	favoList, err := api.GetFavorites(nil)
	if err != nil {
		log.Fatal(err)
		fmt.Print("get favolist is failed")
	}

	// いいねするツイートのスライス(箱)
	var tweets []Tweet

	// 1回の処理でいいねするツイートの数
	favoCount := config.Config.FavoCount

	// 5回分のrandを入れるスライス作成
	var target []int

	// 後で使われるrandの定義(for文の外で定義しないと整合性が保たれない)
	rand.Seed(time.Now().UnixNano())

	// ツイートのバリデーションをかける
ROOP:
	for _, data := range searchTweet.Statuses {

		// 「プロサー」アカウントを除く
		if data.User.Name == "プロサー" {
			continue ROOP
		}

		// 既にいいねしてるツイートを除く
		for _, favoTweet := range favoList {
			if favoTweet.Id == data.Id {
				continue ROOP
			}
		}

		// log用にツイート情報をまとめる
		tweet := Tweet{}
		tweet.Id = data.Id
		tweet.Text = data.Text
		tweet.User = data.User.Name
		tweet.CreatedAt = data.CreatedAt
		tweets = append(tweets, tweet)
	}

	// いいねを押すツイートを決定するためのrandのスライスを生成する
	// 今の仕様では数字が被る回数がランダムなのでいいね回数も変則的になる
	for i := 0; i < favoCount; i++ {
		rand := rand.Intn(len(tweets))
		result := model.IntContains(target, rand)
		if result == false {
			target = append(target, rand)
		}
	}

	// 「いいね」を押しにいく
	for index, tweet := range tweets {
		result := model.IntContains(target, index)
		if result == true {
			newFavo, err := api.Favorite(tweet.Id)
			if err != nil {
				log.Println("favorite tweet is failed")
			}
			log.Println("お気に入りしたツイート")
			log.Println(newFavo)
		}
	}
}

type Tweet struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
}
