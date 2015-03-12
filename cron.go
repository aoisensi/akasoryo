package main

import (
	crontab "github.com/robfig/cron"
)

var (
	cron *crontab.Cron
)

func initCron() error {
	cron = crontab.New()
	cron.AddFunc("0 0 0 * * *", func() {})
	cron.Start()
	return nil
}

func changedDay() {
	t := nowJST()
	weekday := "日月火水木金土"[int(t.Weekday())]
	tweet(textf("CHANGED_DAY_TWEET", master, t.Year(), int(t.Month()), t.Day(), weekday))
}
