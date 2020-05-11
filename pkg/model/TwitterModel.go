package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func ConnectTwitterApi() (*anaconda.TwitterApi, error) {

	// 構造体にセット
	twitterAccount := TwitterAccount{}
	twitterAccount.AccessToken = os.Getenv("ACCESS_TOKEN")
	twitterAccount.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	twitterAccount.ConsumerKey = os.Getenv("CONSUMER_KEY")
	twitterAccount.ConsumerSecret = os.Getenv("CONSUMER_KEY_SECRET")

	fmt.Print(twitterAccount)

	return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret), nil
}

func AccountInfo() (map[string]string, error) {
	row, err := ioutil.ReadFile("path/to/twitterAccount.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var twitterAccount TwitterAccount

	// 構造体にセット
	json.Unmarshal(row, &twitterAccount)
	return map[string]string{
		"AccessToken":       twitterAccount.AccessToken,
		"AccessTokenSecret": twitterAccount.AccessTokenSecret,
		"ConsumerKey":       twitterAccount.ConsumerKey,
		"ConsumerSecret":    twitterAccount.ConsumerSecret,
	}, nil
}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
