package main

import (
	"fmt"
	"testing"
)

func InMenu(id int, menu map[int]string) bool {
	for k := range menu {
		if k == id {
			return true
		}
	}
	return false
}

func TestNoon_Random(t *testing.T) {
	fmt.Println("menu:", noon.Len())
	for i := 0; i < 100; i++ {
		index := noon.Random()
		if b := InMenu(index, noon.Menu); !b {
			t.Errorf("menu index: %d not exists", index)
		}
	}
}

func TestNoon_HaveLen(t *testing.T) {
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	if noon.HaveLen() != 5 {
		t.Errorf("len: %d", noon.HaveLen())
	}
	noon.Result()
	if noon.HaveLen() != 0 {
		t.Errorf("len: %d", noon.HaveLen())
	}
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	if noon.HaveLen() != 5 {
		t.Errorf("len: %d", noon.HaveLen())
	}
}
func TestNoon_History(t *testing.T) {
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	fmt.Println(noon.History())
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	noon.Result()
	fmt.Println(noon.History())
}
