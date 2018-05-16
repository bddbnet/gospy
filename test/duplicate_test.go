package test

import (
	"fmt"
	"testing"

	"github.com/bddbnet/gospy/engine"
)

func TestDuplicate(t *testing.T) {
	b := engine.IsDuplicate("https://www.baidu.com/1")
	fmt.Println(b)
}
