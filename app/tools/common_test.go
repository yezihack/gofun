package tools

import "testing"

func TestCommon_CheckIsWeek(t *testing.T) {
	c := new(Common)
	b := c.CheckIsWeek()
	if b != true {
		t.Error(b)
	}
}
