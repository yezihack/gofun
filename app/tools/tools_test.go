package tools

import (
	"fmt"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	f := GetCurrentDirectory()
	fmt.Println(f)
}
