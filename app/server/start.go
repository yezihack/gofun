package server

import (
	"fmt"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/tools"
)

func Start(c *cron.Cron) {
	ding := new(tools.DingDing)
	office := Office{}

	Meal.Fix(Config.Fix...)
	fmt.Println(Meal.History())

	//随机吃饭
	c.AddFunc("0 30 11 * * 1-5", func() {
		meal := Meal.Random()
		k3log.Info("随机吃饭", meal)
		err := ding.Send(Config.Token.Meal, ding.SetText(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	c.AddFunc("0 0 12 * * 5", func() {
		historyWeek := Meal.History() //历史回顾
		k3log.Info("一周回顾", historyWeek)
		err := ding.Send(Config.Token.Meal, ding.SetText(historyWeek, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//下班
	c.AddFunc("0 30 18 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 45 18 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 0 19 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//上班
	c.AddFunc("0 0 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 15 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.AddFunc("0 25 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	c.Start()
}
