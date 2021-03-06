package handler

import (
	"GoODer/config"
	"GoODer/pkg/model"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func CreateFavoriteHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 認証
		api, err := model.ConnectTwitterApi()
		if err != nil {
			log.Fatal("could not connect API")
		}
		log.Print("API connect is success")

		// 指定語句のツイートをGET(15件？)
		procirList, err := api.GetListTweets(1260478553738694658, false, nil)
		if err != nil {
			log.Fatal("get list is failed")
		}
		/*searchTweet, err := api.GetSearch(config.Config.FavoWord, nil)
		if err != nil {
			log.Fatal(err)
			fmt.Print("search tweet is failed")
		}*/

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
		//for _, data := range searchTweet.Statuses {
		for _, data := range procirList {

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
				_, err := api.Favorite(tweet.Id)
				if err != nil {
					log.Println("favorite tweet is failed")
				}
			}
		}
	}
}

func AutoFollowHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// api認証
		api, err := model.ConnectTwitterApi()
		if err != nil {
			log.Fatal("could not connect API")
		}

		// フォローするアカウントの抽出
		accounts, err := api.GetUserSearch(config.Config.FollowWord, nil)
		if err != nil {
			log.Fatal("get follow account is failed")
		}

		// フォローするアカウントを決めるための乱数生成
		rand.Seed(time.Now().UnixNano())
		var target []int
		for i := 0; i < config.Config.FollowCount; i++ {
			rand := rand.Intn(len(accounts))
			target = append(target, rand)
		}

		// 乱数と照らし合わせフォローアカウントを確定
	ROOP:
		for index, account := range accounts {
			if account.Following == true {
				log.Printf("%s is still following", account.Name)
				continue ROOP
			}
			result := model.IntContains(target, index)
			if result == true {
				_, err := api.FollowUserId(account.Id, nil)
				if err != nil {
					log.Println("follow user is failed")
				} else {
					log.Printf("follow %s is success", account.Name)
				}
			}
		}
	}
}

type Tweet struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
}
