package tools

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGoCache_Get(t *testing.T) {
	tc := New(DefaultExpiration)
	val, err := tc.Get("aa")
	if err == nil || val != nil {
		t.Error("键不存在")
	}
	tc.Put("aa", 100, 10*time.Millisecond)
	<-time.After(5 * time.Millisecond)
	val, err = tc.Get("aa")
	if err != nil {
		t.Error(err)
	} else if val == nil {
		t.Error("aa val is null")
	}

	tc.Put("bb", 500, 20*time.Millisecond)
	<-time.After(5 * time.Millisecond)
	val, err = tc.Get("bb")
	if err != nil {
		t.Error(err)
	} else if val == nil {
		t.Error("bb val is null")
	}

	tc.Put("cc", "goCache", DefaultExpiration)
	<-time.After(100 * time.Millisecond)
	val, err = tc.Get("cc")
	fmt.Println(val)
	if err != nil {
		t.Error("cc", err)
	} else if val == nil {
		t.Error("cc val is null")
	} else if a := val.(string); !strings.EqualFold(a, "goCache") {
		t.Error("cc value not equal to goCache")
	}
}

func TestGoCache_PutDefault(t *testing.T) {
	tc := New(time.Millisecond * 500)
	tc.PutDefault("aa", 1)
	<-time.After(510 * time.Millisecond)
	val, err := tc.Get("aa")
	if err == nil {
		t.Error("aa err is not nil")
	} else if val != nil {
		t.Error("aa val is not nil", val)
	}

	tc.PutDefault("bb", "3king")
	val, err = tc.Get("bb")
	<-time.After(400 * time.Millisecond)
	if err != nil {
		t.Error("bb", err)
	} else if val == nil {
		t.Error("bb value is null")
	} else if a := val.(string); a != "3king" {
		t.Error("bb value is not equal to 3king")
	}

}
