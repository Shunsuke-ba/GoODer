package model

import (
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

	return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret), nil
}

// 配列に文字列があるかの関数 ある→true ない→false
func StringContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// 配列に数宇があるかの関数 ある→true ない→false
func IntContains(arr []int, str int) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

/*func AccountInfo() (map[string]string, error) {
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
}*/
