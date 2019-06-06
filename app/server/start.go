package server

import (
	"fmt"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/tools"
)

func Start(c *cron.Cron, stopChan chan struct{}) (err error) {
	ding := new(tools.DingDing)
	office := Office{}
	Config := Serve.Config

	Meal.Fix(Config.Fix...)
	fmt.Println(Meal.History())

	i := 0
	for i > 10 {
		i ++
		food := Meal.Random()
		fmt.Println(food)
	}

	return
	//随机吃饭
	err = c.AddFunc("0 30 11 * * 1-5", func() {
		meal := Meal.Random()
		k3log.Info("随机吃饭", meal)
		err := ding.Send(Config.Token.Meal, ding.SetText(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	err = c.AddFunc("0 0 12 * * 5", func() {
		historyWeek := Meal.History() //历史回顾
		k3log.Info("一周回顾", historyWeek)
		err := ding.Send(Config.Token.Meal, ding.SetText(historyWeek, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//下班
	err = c.AddFunc("0 30 18 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 45 18 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 0 19 * * 1-5", func() {
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//上班
	err = c.AddFunc("0 0 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 15 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 25 9 * * 1-5", func() {
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//开启
	c.Start()
	//关闭
	go func() {
		<-stopChan
		c.Stop()
	}()
	return
}
