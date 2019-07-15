package tools

import (
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

//判断文件是否存在
func CheckFileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
