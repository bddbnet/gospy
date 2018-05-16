package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"LearnGo/awe/spy2/parser/h.bilibili.com"
)

func TestPage(t *testing.T) {
	fmt.Println(h_bilibili_com.Page(30, 99)) //1
	fmt.Println(h_bilibili_com.Page(30, 3))  //1
	fmt.Println(h_bilibili_com.Page(30, 11)) //1

	fmt.Println(time.Millisecond * time.Duration(RandInt(1000, 3000)))

}

func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
