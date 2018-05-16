package parser_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/bddbnet/gospy/engine"
	"github.com/bddbnet/gospy/model"
	"github.com/bddbnet/gospy/parser/h.bilibili.com"
)

// step 4 获取用户图片总数
func TestUserUploadCount(t *testing.T) {
	bytes, err := ioutil.ReadFile("count.json")
	if err != nil {
		t.Error(err)
	}

	j := engine.UploadCount{}
	json.Unmarshal(bytes, &j)
	userUploadCount := model.UserUploadCount{}
	userUploadCount.DllCount = j.Data.AllCount
	userUploadCount.DailyCount = j.Data.DailyCount
	userUploadCount.DrawCount = j.Data.DrawCount
	userUploadCount.PhotoCount = j.Data.PhotoCount

	parseResult := h_bilibili_com.UserUploadCount(bytes, "")

	if userUploadCount != parseResult.Items[0].Payload {
		t.Error("not match")
		fmt.Println(userUploadCount)
		fmt.Println("----------")
		fmt.Println(parseResult.Items[0])
	}

}
