package main

import (
	"fmt"
	"log"
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
		envs[i] = os.Getenv(v)
		if envs[i] == "" {
			return fmt.Errorf("Enviroment \"%s\" is not set.", v)
		}
		log.Printf("Set env \"%v\"", v)
	}
	anaconda.SetConsumerKey(envs[0])
	anaconda.SetConsumerSecret(envs[1])
	twitter = anaconda.NewTwitterApi(envs[2], envs[3])
	return nil
}

func tweet(text string) {
	_, err := twitter.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
}
