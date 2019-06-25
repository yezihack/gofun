package server

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/tools"
)

type Office struct {
}

//上班
func (Office) On() string {
	s := fmt.Sprintf(config.Template["office_on"], Serve.Config.Office.On, tools.GetNow().Format("15:04:05"))
	diff := tools.SubMinute(Serve.Config.Office.On)
	if diff > 0 {
		return s + ", 还有" + strconv.Itoa(diff) + "分钟"
	} else {
		f, _ := strconv.ParseFloat(strconv.Itoa(diff), 10)
		return s + ", 已过" + fmt.Sprint(math.Abs(f)) + "分钟, 您已经迟到啦"
	}
}

//下班
func (Office) Off() string {
	s := fmt.Sprintf(config.Template["office_off"], Serve.Config.Office.Off, tools.GetNow().Format("15:04:05"))
	diff := tools.SubMinute(Serve.Config.Office.Off)
	if diff > 0 {
		return s + ", 还有" + strconv.Itoa(diff) + "分钟"
	} else {
		f, _ := strconv.ParseFloat(strconv.Itoa(diff), 10)
		s += ", 已过" + fmt.Sprint(math.Abs(f)) + "分钟, 别忘记打卡啦"
		if math.Abs(f) > 10 && math.Abs(f) < 60 {
			s += ",工作诚可贵,身体价更高,下班回家陪家人"
		} else if math.Abs(f) > 60 {
			s += ",老铁还在加班呢.辛苦啦"
		}
		return s
	}
}
