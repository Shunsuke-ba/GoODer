package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func ConnectTwitterApi() (*anaconda.TwitterApi, error) {
	row, err := ioutil.ReadFile("path/to/twitterAccount.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var twitterAccount TwitterAccount
	// 構造体にセット
	json.Unmarshal(row, &twitterAccount)

	return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret), nil
}

func (twitterAccount TwitterAccount) AccountInfo() (map[string]string, error) {
	row, err := ioutil.ReadFile("path/to/twitterAccount.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

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
