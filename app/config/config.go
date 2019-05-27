package config

import "time"

var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

//钉钉URL
var DingDingUrl = map[string]string{
	"meal":   "https://oapi.dingtalk.com/robot/send?access_token=75059ec45d06760e3addac0298300751ab761d9df781c057fcb47bb03040f644", //吃饭
	"office": "https://oapi.dingtalk.com/robot/send?access_token=7ead7b8070e828e9ded79f80570dc65354e2aaece1624f80cc20e8222c06e699", //上下班提醒
}

//餐厅列表
var DiningRoot = map[int]string{
	1: "地下城",
	2: "二宝",
	3: "麻辣汤",
	4: "拉面",
	5: "第二家",
	6: "牛肉面",
	7: "开发下家",
}

//模板
var Template = map[string]string{
	"office_on":  "老铁上班打卡啦,距离9:30,现在:%s",
	"office_off": "老铁下班打卡啦,下班18:30,现在:%s",
}
