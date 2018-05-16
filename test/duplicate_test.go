package test

import (
	"LearnGo/awe/spy2/engine"
	"fmt"
	"testing"
)

func TestDuplicate(t *testing.T) {
	b := engine.IsDuplicate("https://www.baidu.com/1")
	fmt.Println(b)
}
