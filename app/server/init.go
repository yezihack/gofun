package server

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/tools"
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
	Token string `toml:"token"`
}

//服务配置结构体
type Servers struct {
	Config     Conf   //配置信息
	ConfigName string //配置文件名称
	ConfigPath string //配置文件全路径 ,带文件名
	execPath   string //exec执行目录
}

var Serve Servers

//初使配置文件
func init() {
	Serve.ConfigName = config.ConfName
	Serve.execPath = tools.GetCurrentDirectory() + "/"
	Serve.ConfigPath = Serve.execPath + Serve.ConfigName
	//parse params
	Serve.flagParse()
	//load config file
	Serve.LoadConfig()
}

//加载配置文件
func (s *Servers) LoadConfig() bool {
	//解析toml
	if _, err := toml.DecodeFile(Serve.ConfigPath, &Serve.Config); err != nil {
		k3log.Error("initConf", err)
		return false
	}
	if Serve.Config.Token.Token == "" ||
		len(Serve.Config.Meal.List) == 0 ||
		Serve.Config.Office.Off == "" ||
		Serve.Config.Office.On == "" {
		k3log.Warn("配置文件为空,请填写有效信息")
		return false
	}
	return true
}

//处理运行的参数数据
func (s *Servers) flagParse() {
	var (
		configName, fix string
	)
	flag.StringVar(&configName, "c", "", "初使项目")
	flag.StringVar(&fix, "fix", "", "修复数据")
	flag.Parse()

	//自定义配置名称
	if configName != "" {
		Serve.ConfigName = configName
		file := Serve.execPath + Serve.ConfigName
		if !tools.CheckFileExists(file) {
			s.WriteConfig()
		}
		os.Exit(0)
	}
	//处理传过来的数据
	if !strings.EqualFold(fix, "") {
		fixInt := make([]int, 0)
		for _, val := range strings.Split(fix, ",") {
			v, _ := strconv.Atoi(val)
			fixInt = append(fixInt, v)
		}
		Serve.Config.Fix = fixInt
	}
}

//配置模板输出
func (*Servers) WriteConfig() {
	var data bytes.Buffer
	data.WriteString("title = \"gofun的娱乐版\"")
	data.WriteString("\n")
	data.WriteString("\n")
	data.WriteString("[access_token]")
	data.WriteString("\n")
	data.WriteString("token = \"\"")
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
	err := ioutil.WriteFile(path+Serve.ConfigName, data.Bytes(), 0666)
	if err != nil {
		k3log.Error("WriteConfig", err)
		os.Exit(0)
	}
}
