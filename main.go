package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var master string

func init() {
	master = os.Getenv("MASTER")
	inits := []func() error{
		initTwitter,
		initHTTP,
		initCron,
	}
	for _, f := range inits {
		err := f()
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	}
}

func main() {
	cron.Start()
	launchedServer()
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func launchedServer() {
	t := time.Now().UTC()
	tweet(textf("LAUNCHED_TWEET", t, master))
}

func getText(key string) string {
	v, ok := texts[key]
	if ok {
		return v
	}
	return fmt.Sprintf("ERROR NO TEXT DATA \"%s\"", key)
}

func textf(key string, a ...interface{}) string {
	return fmt.Sprintf(getText(key), a...)
}

func nowJST() time.Time {
	l, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(l)
}
