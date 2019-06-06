package server

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/yezihack/gofun/app/config"
)

type MealInfo struct {
	Index int //菜单下标
	Food string //食物
	Week string //何星期
}

type MealStructure struct {
	HaveMeal  []MealInfo     //已使用
	Doing []int
	Menu      map[int]string //菜单
	Config    Conf
}

var Meal MealStructure

func init() {
	Meal.Config  = Serve.Config
	Meal.Init()
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

//随机一个值
func (ms *MealStructure) Random() string {
	RandomContinue:
	//随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if ms.HistoryLen() == 5 { //满五则重置
		ms.Reset()
	}
	//随机一个数
	index := r.Intn(ms.Len())
	//判断随机的值是否已经使用过
	if b := ms.InArray(index); b {
		goto RandomContinue
	}
	//获取对应的值
	food := ms.Menu[index]
	//随机过的存储起来
	ms.HaveMeal = append(ms.HaveMeal, MealInfo{
		Index:index,
		Food: food,
		Week:ms.Week(),
	})
	return food
}

//修复数据
func (ms *MealStructure) Fix(req ...int) {
	for _, idx := range req {
		ms.HaveMeal = append(ms.HaveMeal, MealInfo{
			Index:idx,
			Food: ms.Menu[idx],
		})
	}
	fmt.Println(ms.HaveMeal)
}
