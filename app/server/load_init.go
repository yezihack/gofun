package server

import (
	"github.com/BurntSushi/toml"
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/tools"
	"os"
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
}

//token结构体
type TokenConf struct {
	Meal   string `toml:"meal"`
	Office string `toml:"office"`
}

var Config Conf

//初使配置文件
func init() {
	path := tools.GetCurrentDirectory()
	file := path + "/config.toml"
	if _, err := toml.DecodeFile(file, &Config); err != nil {
		k3log.Error("initConf", err)
		os.Exit(400)
	}
}
