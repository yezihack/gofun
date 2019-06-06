package config

import "time"

var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

//http://tool.bitefu.net/jiari/?d=2019-06-02&back=json&info=1
//日历信息
const SolarUrl = "http://tool.bitefu.net/jiari/?d=%s&back=json&info=1"

//模板
var Template = map[string]string{
	"office_on":  "老铁上班打卡啦,上班时间%s,现在%s",
	"office_off": "老铁下班打卡啦,下班时间%s,现在%s",
}

//配置名称
const ConfName = "gofun.toml"

