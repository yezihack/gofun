package logcolor

import (
	"testing"
	"os"
	"bytes"
	"fmt"
)

func TestDebug(t *testing.T) {
	Info("      饮酒·其五     ")
	Info("              -----%s", "陶渊明")
	Debug("结庐在人境，而无车马喧。")
	Info("问君何能尔？心远地自偏。")
	Warning("采菊东篱下，悠然见南山。")
	Error("山气日夕佳，飞鸟相与还。")
	Debug("此中有真意，欲辨已忘言。")
}
func TestLogger_SetLevel(t *testing.T) {
	lg := New(os.Stdout, true)
	lg.SetLevel(INFO)
	lg.SetColor(false)
	lg.Info("人生如逆旅，我亦是行人-------------%s", "出自宋代苏轼")
	lg.Debug("输不出来的")
	lg.SetLevel(DEBUG)
	lg.Debug("我胡汉三又回来了")
}

func TestNew(t *testing.T) {
	//把结果输入到buffer里
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	lg.Warning("人间正道是沧桑")
	//输出结果
	fmt.Print(buff.String())
}