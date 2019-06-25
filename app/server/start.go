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

	//i := 0
	//for i < 21 {
	//	i++
	//	Meal.Result()
	//	fmt.Println(meal)
	//}
	//fmt.Println(Meal.IsWeek())
	////os.Exit(0)
	//return

	//随机吃饭
	err = c.AddFunc("0 30 11 * * *", func() {
		if !Meal.IsWeek() {
			return
		}
		meal := Meal.Result()
		k3log.Info("随机吃饭", meal)
		err := ding.Send(Config.Token.Token, ding.SetText(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	err = c.AddFunc("0 0 12 * * 5", func() {
		historyWeek := Meal.History() //历史回顾
		k3log.Info("一周回顾", historyWeek)
		err := ding.Send(Config.Token.Token, ding.SetText(historyWeek, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//下班
	err = c.AddFunc("0 30 18 * * *", func() {
		if !Meal.IsWeek() {
			return
		}
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Token, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 45 18 * * *", func() {
		if !Meal.IsWeek() {
			return
		}
		data := office.Off()
		k3log.Info("下班", data)
		err := ding.Send(Config.Token.Token, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//上班
	err = c.AddFunc("0 25 9 * * *", func() {
		if !Meal.IsWeek() {
			return
		}
		data := office.On()
		k3log.Info("上班", data)
		err := ding.Send(Config.Token.Token, ding.SetText(data, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	err = c.AddFunc("0 0 8 * * *", func() {
		data := ding.C.SuanGua()
		s := fmt.Sprintf("今天是%s\n老黄历:\n农历:%s 节气:%s \n宜: %s\n忌: %s\n----今天星期%s是一年中的第%s天",
			data.TypeName, data.NongLiCn, data.JieQi, data.Suit, data.Avoid, data.WeekCn, data.DayNum)
		err := ding.Send(Config.Token.Token, ding.SetText(s, false))
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
