package tools

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeKing2018/k3log"
	"github.com/mikemintang/go-curl"
	"github.com/pkg/errors"
	"github.com/yezihack/gofun/app/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DingDing struct {
}

//发送消息
func (d *DingDing) Send(token string, data map[string]interface{}) (err error) {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
	header := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetHeaders(header).
		SetUrl(url).
		SetPostData(data).
		Post()
	if err != nil {
		return
	}
	reply := resp.Body
	type bd struct {
		ErrorCode int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}
	msg := bd{}
	err = json.Unmarshal([]byte(reply), &msg)
	if err != nil {
		return
	}
	if msg.ErrorCode == 0 {
		return
	}
	err = errors.New(msg.ErrMsg)
	return
}

//设置钉钉消息体
func (d *DingDing) SetText(message string, isAtAll bool) map[string]interface{} {
	reply := make(map[string]interface{})
	reply["msgtype"] = "text"
	reply["text"] = map[string]string{
		"content": message,
	}
	reply["at"] = map[string]bool{
		"isAtAll": isAtAll,
	}
	return reply
}

func GetNow() time.Time {
	return time.Now().In(config.BeijingLocation)
}

//计算时间差, 正则相离, 负则已过
func SubMinute(HourMinute string) int {
	Ymd := "2006-01-02"
	YmdAll := "2006-01-02 15:04:05"
	now := time.Now()
	start := now.Format(Ymd) + " " + HourMinute + ":00"
	s, err := time.ParseInLocation(YmdAll, start, config.BeijingLocation)
	if err != nil {
		return 0
	}
	return int(s.Sub(now.In(config.BeijingLocation)).Minutes())
}

//读取运行的目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		k3log.Error("读取目录", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
