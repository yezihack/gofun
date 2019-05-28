package server

import (
	"fmt"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/tools"
	"math"
	"strconv"
)

type Office struct {
}

//上班
func (Office) On() string {
	s := fmt.Sprintf(config.Template["office_on"], Config.Office.On, tools.GetNow().Format("15:04:05"))
	diff := tools.SubMinute(Config.Office.On)
	if diff > 0 {
		return s + ", 还有" + strconv.Itoa(diff) + "分钟"
	} else {
		f, _ := strconv.ParseFloat(strconv.Itoa(diff), 10)
		return s + ", 已过" + fmt.Sprint(math.Abs(f)) + "分钟, 您已经迟到啦"
	}
}

//下班
func (Office) Off() string {
	s := fmt.Sprintf(config.Template["office_off"], Config.Office.Off, tools.GetNow().Format("15:04:05"))
	diff := tools.SubMinute(Config.Office.Off)
	if diff > 0 {
		return s + ", 还有" + strconv.Itoa(diff) + "分钟"
	} else {
		f, _ := strconv.ParseFloat(strconv.Itoa(diff), 10)
		s += ", 已过" + fmt.Sprint(math.Abs(f)) + "分钟, 别忘记打卡啦"
		if math.Abs(f) > 30 {
			s += ",老兄还不下班,下班都半小时啦"
		} else if math.Abs(f) > 60 {
			s += ",老铁还在加班呢.辛苦啦"
		}
		return s
	}
}
