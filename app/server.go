package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

//已经使用过的清单
type EatList struct {
	Index int    //菜单下标
	Food  string //食物
	Week  string //何星期
}

//午餐结构
type Noon struct {
	HaveMeal []EatList //已使用
	Doing    []int
	Menu     map[int]string //菜单
}

var noon Noon

//初使化
func init() {
	noon.Init()
}

//初使化
func (ms *Noon) Init() {
	ms.Menu = make(map[int]string)
	//加载菜单
	for id, val := range NoonRootConf {
		ms.Menu[id] = val
	}
}

//菜单长度
func (ms *Noon) Len() int {
	return len(ms.Menu)
}

//已经存在的长度
func (ms *Noon) HaveLen() int {
	return len(ms.HaveMeal)
}

//重置
func (ms *Noon) Reset() {
	fmt.Println("Reset", ms.HaveMeal)
	ms.HaveMeal = append(ms.HaveMeal[:0])
	fmt.Println("Reset", ms.HaveMeal)
}

//查看历史记录
func (ms *Noon) History() string {
	his := bytes.Buffer{}
	his.WriteString("本周:\n")
	for i, item := range ms.HaveMeal {
		his.WriteString(fmt.Sprintf("星期%s 吃了<%s>\n", ms.WeekChina(i+1), item.Food))
	}
	return his.String()
}

//星期配置
func (ms *Noon) WeekChina(i int) string {
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
func (ms *Noon) Week() string {
	return time.Now().Weekday().String()
}

//判断是否存在
func (ms *Noon) InArray(id int) bool {
	for _, item := range ms.HaveMeal {
		if item.Index == id {
			return true
		}
	}
	return false
}

//获取结果
func (ms *Noon) Result() string {
	if ms.HaveLen() == 5 { //满五则重置
		fmt.Println("重置了")
		ms.Reset()
	}
	//获取对应的值
	index := ms.Random()
	food := ms.Menu[index]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(21) //每月开发一家
	if num == 1 {
		index = -1
		food = "开发下一家"
	}
	//随机过的存储起来
	ms.HaveMeal = append(ms.HaveMeal, EatList{
		Index: index,
		Food:  food,
		Week:  ms.Week(),
	})
	fmt.Println("Result", ms.HaveMeal)
	return fmt.Sprintf("今日午餐: %s, <备选:%s>", food, ms.Menu[ms.Random()])
}

//随机一个值
func (ms *Noon) Random() int {
RandomContinue:
	//随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//随机一个数
	index := r.Intn(ms.Len())
	//判断随机的值是否已经使用过
	if b := ms.InArray(index); b {
		goto RandomContinue
	}
	return index
}
