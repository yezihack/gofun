package server

import (
	"bytes"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/tools"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//定义配置结构体
type Conf struct {
	Title string    `toml:"title"`
	Token TokenConf `toml:"access_token"`
	Meal  struct {
		List []string `toml:"list"`
	} `toml:"dining_room"`
	Office struct {
		On  string `toml:"on"`
		Off string `toml:"off"`
	} `toml:"office"`
	Fix []int
}

//token结构体
type TokenConf struct {
	Meal   string `toml:"meal"`
	Office string `toml:"office"`
}

var Config Conf
var ConfigFile = "/gofun.toml"

//初使配置文件
func init() {
	path := tools.GetCurrentDirectory()
	file := path + ConfigFile
	var (
		op, fix string
	)
	flag.StringVar(&op, "c", "", "初使项目")
	flag.StringVar(&fix, "fix", "", "修复数据")
	flag.Parse()

	if op != "" {
		ConfigFile = "/" + op
		file = path + ConfigFile
		if !tools.CheckFileExists(file) {
			WriteConfig()
		}
		os.Exit(0)
	}
	if !strings.EqualFold(fix, "") {
		fixInt := make([]int, 0)
		for _, val := range strings.Split(fix, ",") {
			v, _ := strconv.Atoi(val)
			fixInt = append(fixInt, v)
		}
		Config.Fix = fixInt
	}
	//解析toml
	if _, err := toml.DecodeFile(file, &Config); err != nil {
		k3log.Error("initConf", err)
		os.Exit(400)
	}
	if Config.Token.Office == "" ||
		Config.Token.Meal == "" ||
		len(Config.Meal.List) == 0 ||
		Config.Office.Off == "" ||
		Config.Office.On == "" {
		k3log.Warn("配置文件为空,请填写有效信息")
		os.Exit(0)
	}

}

func WriteConfig() {
	var data bytes.Buffer
	data.WriteString("title = \"gofun的娱乐版\"")
	data.WriteString("\n")
	data.WriteString("\n")
	data.WriteString("[access_token]")
	data.WriteString("\n")
	data.WriteString("meal = \"\"")
	data.WriteString("\n")
	data.WriteString("office = \"\"")
	data.WriteString("\n")
	data.WriteString("\n")
	data.WriteString("#餐厅列表")
	data.WriteString("\n")
	data.WriteString("[dining_room]")
	data.WriteString("\n")
	data.WriteString("list = [")
	data.WriteString("\n")
	data.WriteString("]")
	data.WriteString("\n")
	data.WriteString("\n")
	data.WriteString("#上班")
	data.WriteString("\n")
	data.WriteString("[office]")
	data.WriteString("\n")
	data.WriteString("on = \"09:30\" # 上班时间")
	data.WriteString("\n")
	data.WriteString("off = \"18:30\" # 下班时间")
	path := tools.GetCurrentDirectory()
	err := ioutil.WriteFile(path+ConfigFile, data.Bytes(), 0666)
	if err != nil {
		k3log.Error(err)
		os.Exit(0)
	}
}
