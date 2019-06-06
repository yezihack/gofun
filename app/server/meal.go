package server

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	Meal.HaveMeal = make([]MealInfo, 0)
	Meal.Menu = make(map[int]string)
	//加载菜单
	for id, val := range Meal.Config.Meal.List {
		id++
		Meal.Menu[id] = val
	}
}

//动态计算正在执行的
func (ms *MealStructure) CalcDoing() {
	for idx := range Meal.Menu {
		for _, item := range Meal.HaveMeal {
			if item.Index != idx {
				Meal.Doing = append(Meal.Doing, idx)
			}
		}
	}
}

func (ms *MealStructure) Len() int {
	return len(Meal.Doing)
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

//随机一个值
func (ms *MealStructure) Random() string {
	//随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if ms.HistoryLen() == 5 { //满五则重置
		ms.Reset()
	}
	ms.CalcDoing()
	//随机一个数
	index := r.Intn(ms.Len())
	//获取对应的值
	mealKey := ms.DoingMeal[index]
	//将随机的值从切片里删除
	ms.DoingMeal = append(ms.DoingMeal[:index], ms.DoingMeal[index+1:]...)
	//获取名称
	food := Meal.Menu[mealKey]
	//随机过的存储起来
	ms.HaveMeal = append(ms.HaveMeal, MealInfo{
		Food: food,
	})
	return "今天随机:" + food
}

//修复数据
func (ms *MealStructure) Fix(req ...int) {
	for _, idx := range req {
		ms.HaveMeal = append(ms.HaveMeal, MealInfo{
			Food: ms.Menu[idx],
		})
		for k, v := range ms.DoingMeal {
			if v == idx {
				ms.DoingMeal = append(ms.DoingMeal[:k], ms.DoingMeal[k+1:]...)
			}
		}
	}
	fmt.Println(ms.DoingMeal)
}

//动态加载数据
func (ms *MealStructure) DynamicFix(filePath string) {
	LoadConfig(filePath)
	for id, val := range Config.Meal.List {
		id++
		for _, food := range Meal.HaveMeal {
			if food.Food != val {
				Meal.DoingMeal = append(Meal.DoingMeal, id)
				Meal.Menu[len(Meal.Menu)-1] = val
			}
		}
	}
	spew.Dump(Meal)
}
