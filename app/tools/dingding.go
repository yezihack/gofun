package tools

import (
	"encoding/json"
	"fmt"
	curl "github.com/mikemintang/go-curl"
	"github.com/pkg/errors"
)

type DingDing struct {
	C Common
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
