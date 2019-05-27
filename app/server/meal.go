package server

import (
	"bytes"
	"fmt"
	"github.com/yezihack/gofun/app/config"
	"math/rand"
	"time"
)

type MealInfo struct {
	Week string //星期几
	Food string //食物
}

type MealStructure struct {
	HaveMeal  []MealInfo //已使用
	DoingMeal []int      //未使用
}

var Meal MealStructure

func init() {
	Meal.Init()
}

//初使化
func (ms *MealStructure) Init() {
	Meal.DoingMeal = make([]int, 0)
	Meal.DoingMeal = make([]int, 0)
	for id := range config.DiningRoot {
		Meal.DoingMeal = append(Meal.DoingMeal, id)
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
	for _, item := range ms.HaveMeal {
		his.WriteString(fmt.Sprintf("%s 吃了<%s>\n", item.Week, item.Food))
	}
	return his.String()
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
	value := ms.DoingMeal[index]
	//将随机的值从切片里删除
	ms.DoingMeal = append(ms.DoingMeal[:index], ms.DoingMeal[index+1:]...)
	//获取名称
	food := config.DiningRoot[value]
	//特殊备选
	if value == 7 {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		food += ", 备选:" + config.DiningRoot[ms.DoingMeal[r.Intn(ms.Len())]]
	}
	//随机过的存储起来
	ms.HaveMeal = append(ms.HaveMeal, MealInfo{
		Food: food,
		Week: ms.Week(),
	})
	return food
}
