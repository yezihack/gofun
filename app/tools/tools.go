package tools

import (
	"encoding/json"
	"github.com/mikemintang/go-curl"
	"github.com/pkg/errors"
	"github.com/yezihack/gofun/app/config"
	"time"
)

func DingDingCurlPost(url string, data map[string]interface{}) (err error) {
	header := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	req := curl.NewRequest()
	resp, err := req.SetHeaders(header).SetUrl(url).SetPostData(data).Post()
	if err != nil {
		return
	}
	reply := resp.Body
	type msgS struct {
		ErrorCode int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}
	msg := msgS{}
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
func SetDingDingData(message string, isAtAll bool) map[string]interface{} {
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
