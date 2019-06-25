package tools

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/config"
	"io/ioutil"
	"net/http"
	"time"
)

type Common struct {
}

//判断是否是工作日
func (c *Common) CheckIsWeek() bool {
	url := "http://tool.bitefu.net/jiari/?d=%s&back=json"
	Ymd := time.Now().Format("2006-01-02")
	url = fmt.Sprintf(url, Ymd)
	resp, err := http.Get(url)
	if err != nil {
		k3log.Error("CheckIsWeek", err)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		k3log.Error("CheckIsWeek", err)
		return false
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		k3log.Error("CheckIsWeek", err)
		return false
	}
	if fmt.Sprint(m[Ymd]) == "0" {
		return true
	}
	return false
}

//
func (c *Common) SuanGua() (result *config.SuanGuaStruct) {
	url := "http://tool.bitefu.net/jiari/?d=%s&back=json&info=1"
	Ymd := time.Now().Format("2006-01-02")
	url = fmt.Sprintf(url, Ymd)
	resp, err := http.Get(url)
	if err != nil {
		k3log.Error("SuanGua", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		k3log.Error("SuanGua", err)
		return
	}
	result = new(config.SuanGuaStruct)
	err = json.Unmarshal(body, &result)
	if err != nil {
		k3log.Error("SuanGua", err)
		return
	}
	return
}
