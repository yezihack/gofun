package server

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ThreeKing2018/goCache"
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/tools"
	"math/rand"
)

type MealInfo struct {
	Index int    //菜单下标
	Food  string //食物
	Week  string //何星期
}

type MealStructure struct {
	HaveMeal []MealInfo //已使用
	Doing    []int
	Menu     map[int]string //菜单
	Config   Conf
	Common   tools.Common
	Cache    goCache.GoCacher
}

var Meal MealStructure

func init() {
	Meal.Config = Serve.Config
	Meal.Init()
	Meal.Cache = goCache.New(time.Hour * time.Duration(23))
}

//初使化
func (ms *MealStructure) Init() {
	ms.HaveMeal = make([]MealInfo, 0)
	ms.Menu = make(map[int]string)
	//加载菜单
	for id, val := range ms.Config.Meal.List {
		ms.Menu[id] = val
	}
}
func (ms *MealStructure) ReInit() {
	Meal.Config = Serve.Config
	ms.Menu = make(map[int]string)
	//加载菜单
	for id, val := range ms.Config.Meal.List {
		ms.Menu[id] = val
	}
	fmt.Println(ms.Menu)
	fmt.Println(ms.HaveMeal)
}

func (ms *MealStructure) Len() int {
	return len(ms.Menu)
}

//历史长度
func (ms *MealStructure) HistoryLen() int {
	return len(ms.HaveMeal)
}

//重置
func (ms *MealStructure) Reset() {
	ms.Init()
}

//查看历史记录
func (ms *MealStructure) History() string {
	his := bytes.Buffer{}
	his.WriteString("本周:\n")
	for i, item := range ms.HaveMeal {
		his.WriteString(fmt.Sprintf("星期%s 吃了<%s>\n", ms.WeekChina(i+1), item.Food))
	}
	return his.String()
}

func (ms *MealStructure) WeekChina(i int) string {
	switch i {
	case 1:
		return "一"
	case 2:
		return "二"
	case 3:
		return "三"
	case 4:
		return "四"
	case 5:
		return "五"
	case 6:
		return "六"
	default:
		return "日"
	}
}

//打印星期
func (ms *MealStructure) Week() string {
	return time.Now().In(config.BeijingLocation).Weekday().String()
}

//判断是否存在
func (ms *MealStructure) InArray(id int) bool {
	for _, item := range ms.HaveMeal {
		if item.Index == id {
			return true
		}
	}
	return false
}

//获取结果
func (ms *MealStructure) Result() string {
	//获取对应的值
	index := ms.Random(true)
	food := ms.Menu[index]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(21) //每月开发一家
	if num == 1 {
		food = "开发下一家"
	}
	//随机过的存储起来
	ms.HaveMeal = append(ms.HaveMeal, MealInfo{
		Index: index,
		Food:  food,
		Week:  ms.Week(),
	})
	return fmt.Sprintf("今日午餐: %s, <备选:%s>", food, ms.Menu[ms.Random(false)])
}

//随机一个值
func (ms *MealStructure) Random(isReset bool) int {
RandomContinue:
	//随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if ms.HistoryLen() == 5 && isReset { //满五则重置
		ms.Reset()
	}
	//随机一个数
	index := r.Intn(ms.Len())
	//判断随机的值是否已经使用过
	if b := ms.InArray(index); b {
		goto RandomContinue
	}
	return index
}

//判断是否是工作日
func (ms *MealStructure) IsWeek() bool {
	val, err := ms.Cache.Get(config.WEEK_KEY)
	if err != nil || val == nil {
		week := ms.Common.CheckIsWeek()
		k3log.Warn("isWeek", "from url")
		ms.Cache.Set(config.WEEK_KEY, week, time.Hour*3)
		return week
	}
	k3log.Info("isWeek", "from cache")
	return val.(bool)
}
