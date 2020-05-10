package handler

import (
	"GoODer/model"
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
			fmt.Print(tweet.Id)
			parseUrl, err := fmt.Printf("/create.json?id=%v", tweet.Id)
			if err != nil {
				log.Fatal("parse url printf is valid")
			}
			referece, err := url.Parse(string(parseUrl))
			if err != nil {
				log.Fatal("query parse is valid")
			}
			endPoint := base.ResolveReference(referece).String()
			req, err := http.NewRequest("POST", endPoint, nil)
			if err != nil {
				log.Fatal("make request is failed")
			}
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
