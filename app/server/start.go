package server

import (
	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/tools"
)

func Start(c *cron.Cron) {
	ding := new(tools.DingDing)
	office := Office{}
	//随机吃饭
	c.AddFunc("0 30 11 * * 1-5", func() {
		meal := Meal.Random()
		err := ding.Send(Config.Token.Meal, ding.SetText(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	c.AddFunc("0 0 12 * * 5", func() {
		historyWeek := Meal.History() //历史回顾
		err := ding.Send(Config.Token.Meal, ding.SetText(historyWeek, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//下班
	c.AddFunc("0 30 18 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.Off(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 45 18 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.Off(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 0 19 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.Off(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//上班
	c.AddFunc("0 0 9 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.On(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 15 9 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.On(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 25 9 * * *", func() {
		err := ding.Send(Config.Token.Office, ding.SetText(office.On(), true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.Start()
}
