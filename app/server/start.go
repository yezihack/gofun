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

	data := ding.C.SuanGua()
	s := fmt.Sprintf("今天是%s\n老黄历: 农历:%s 节气:%s \n宜: %s\n忌: %s\n ----今天星期%s是一年中的第%s天",
		data.TypeName, data.NongLiCn, data.JieQi, data.Suit, data.Avoid, data.WeekCn, data.DayNum)
	fmt.Println(s)

	return
	//随机吃饭
	err = c.AddFunc("0 30 11 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		meal := Meal.Random()
		k3log.Info("随机吃饭", meal)
		err := ding.Send(Config.Token.Meal, ding.SetText(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	err = c.AddFunc("0 0 12 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		historyWeek := Meal.History() //历史回顾
		k3log.Info("一周回顾", historyWeek)
		err := ding.Send(Config.Token.Meal, ding.SetText(historyWeek, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//下班
	err = c.AddFunc("0 30 18 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 45 18 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 0 19 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//上班
	err = c.AddFunc("0 0 9 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 15 9 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 25 9 * * *", func() {
		if !ding.C.CheckIsWeek() {
			return
		}
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Office, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 0 8 * * *", func() {
		data := ding.C.SuanGua()
		s := fmt.Sprintf("老黄历: \n 宜: %s \n  忌:%s", data.Suit, data.Avoid)
		err := ding.Send(Config.Token.Meal, ding.SetText(s, true))
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
