package server

import (
	"bytes"
	"fmt"
	"github.com/yezihack/gofun/app/config"
	"math/rand"
	"time"
)

type MealInfo struct {
	Food string //食物
}

type MealStructure struct {
	HaveMeal  []MealInfo     //已使用
	DoingMeal []int          //未使用
	Menu      map[int]string //菜单
}

var Meal MealStructure

func init() {
	Meal.Init()
}

//初使化
func (ms *MealStructure) Init() {
	Meal.HaveMeal = make([]MealInfo, 0)
	Meal.DoingMeal = make([]int, 0)
	Meal.Menu = make(map[int]string)
	for id, val := range Config.Meal.List {
		id++
		Meal.DoingMeal = append(Meal.DoingMeal, id)
		Meal.Menu[id] = val
	}
}

//长度
func (ms *MealStructure) Len() int {
	return len(Meal.DoingMeal)
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
	default:
		return "五"
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
	if ms.Len() == 0 {
		ms.Reset()
	}
	//随机一个数
	index := r.Intn(ms.Len())
	//获取对应的值
	mealKey := ms.DoingMeal[index]
	//将随机的值从切片里删除
	ms.DoingMeal = append(ms.DoingMeal[:index], ms.DoingMeal[index+1:]...)
	//获取名称
	food := Meal.Menu[mealKey]
	//特殊备选
	if mealKey == 7 {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		food += ", 备选:" + Meal.Menu[ms.DoingMeal[r.Intn(ms.Len())]]
	}
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
