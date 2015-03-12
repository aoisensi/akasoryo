package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var twitter *anaconda.TwitterApi

func initTwitter() error {
	envs := []string{
		"CONSUMER_KEY",
		"CONSUMER_SECRET",
		"ACCESS_TOKEN",
		"ACCESS_TOKEN_SECRET",
	}
	for i, v := range envs {
		if v == "" {
			return fmt.Errorf("Enviroment \"%s\" is not set.", v)
		}
		envs[i] = os.Getenv(v)
	}
	anaconda.SetConsumerKey(envs[0])
	anaconda.SetConsumerSecret(envs[1])
	twitter = anaconda.NewTwitterApi(envs[2], envs[3])
	return nil
}

func tweet(text string) error {
	_, err := twitter.PostTweet(text, nil)
	return err
}