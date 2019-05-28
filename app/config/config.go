package config

import "time"

var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

//模板
var Template = map[string]string{
	"office_on":  "老铁上班打卡啦,上班时间%s,现在%s",
	"office_off": "老铁下班打卡啦,下班时间%s,现在%s",
}
