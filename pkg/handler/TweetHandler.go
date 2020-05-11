package handler

import (
	"GoODer/pkg/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//var accountInfo := model.TwitterAcoount.AccountInfo

func GetTweetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		api, err := model.ConnectTwitterApi()
		if err != nil {
			log.Fatal("could not connect API")
		}
		searchTweet, _ := api.GetSearch(`"#プログラミング"`, nil)

		var tweets []Tweet
	ROOP:
		for _, data := range searchTweet.Statuses {
			if data.User.Name == "プロサー" {
				continue ROOP
			}
			tweet := Tweet{}
			tweet.Id = data.Id
			tweet.Text = data.Text
			tweet.User = data.User.Name
			tweet.Created_at = data.CreatedAt

			tweets = append(tweets, tweet)
		}
		fmt.Fprintln(writer, TweetsResponse{Tweets: tweets})

		// 自動いいね作成

		// Example Request
		/*curl --request POST
		--url 'https://api.twitter.com/1.1/favorites/create.json?id=TWEET_ID_TO_FAVORITE'
		--header 'authorization: OAuth oauth_consumer_key="YOUR_CONSUMER_KEY", oauth_nonce="AUTO_GENERATED_NONCE", oauth_signature="AUTO_GENERATED_SIGNATURE", oauth_signature_method="HMAC-SHA1", oauth_timestamp="AUTO_GENERATED_TIMESTAMP", oauth_token="USERS_ACCESS_TOKEN", oauth_version="1.0"'
		--header 'content-type: application/json'*/

		for _, tweet := range tweets {
			base, err := url.Parse("https://api.twitter.com/1.1/favorites/")
			if err != nil {
				log.Fatal("url.Parse is valid")
			}
			parseUrl, err := fmt.Printf("/create.json?id=%v", tweet.Id)
			if err != nil {
				log.Fatal("parse url printf is valid")
			}
			reference, err := url.Parse(string(parseUrl))
			if err != nil {
				log.Fatal("query parse is valid")
			}
			endPoint := base.ResolveReference(reference).String()
			fmt.Print(endPoint)

			// ランダム文字列でoauth_nonce作成
			/*oauth_nonce, err := uuid.NewRandom()
				if err != nil {
					log.Fatal("rand is failed")
				}

				// timestamp作成
				timestamp := time.Now()

				// signature作るようのパラメータを設定
			/param := fmt.Sprintf("OAuth oauth_consumer_key=%s&oauth_nonce=%s&oauth_signature_method=HMAC-SHA1&oauth_timestamp=%s&oauth_token=%s&oauth_version=1.0", os.Getenv("ConsumerKey"), oauth_nonce, timestamp, os.Getenv("ACCESS_TOKEN"))

				// signature作成
				baseString := fmt.Sprintf("POST&%s&%s", "https://api.twitter.com/1.1/favorites/create.json", values.Encode())
				key := fmt.Sprintf("%s&", auth.SecretID)
				mac := hmac.New(sha1.New, []byte(key))
				mac.Write([]byte(baseString))
				signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))

				// headerを定義
				/*if err != nil {
					log.Fatal("making header is failed")
				}*/

			// リクエスト作成
			req, err := http.NewRequest("POST", endPoint, nil)
			if err != nil {
				log.Fatal("make request is failed")
			}

			// headerをセット
			//req.Header.Set("Authentication", header)
			/*Authorization:
			OAuth oauth_consumer_key="xvz1evFS4wEEPTGEFPHBog",
			oauth_nonce="kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
			oauth_signature="tnnArxj06cWHq44gCs1OSKk%2FjLY%3D",
			oauth_signature_method="HMAC-SHA1",
			oauth_timestamp="1318622958",
			oauth_token="370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
			oauth_version="1.0"*/

			// client周り
			var client = &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("client.do is valid")
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("body read is valid")
			}
			fmt.Println(string(body))
		}
	}
}

type Tweet struct {
	User       string `json:"user"`
	Text       string `json:"text"`
	Created_at string `json:"created_at`
	Id         int64  `json:"id`
}

type TweetsResponse struct {
	Tweets []Tweet `json:"tweets`
}
