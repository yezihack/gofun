package server

import (
	"fmt"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/tools"
)

type Office struct {
}

//上班
func (Office) On() string {
	return fmt.Sprintf(config.Template["office_on"], tools.GetNow().Format("15:04:05"))
}

//下班
func (Office) Off() string {
	return fmt.Sprintf(config.Template["office_off"], tools.GetNow().Format("15:04:05"))
}
