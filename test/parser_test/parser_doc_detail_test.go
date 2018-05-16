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

// step 6 doc 详情

func TestDocDetail(t *testing.T) {

	// 读取测试数据
	bytes, err := ioutil.ReadFile("detail.json")
	if err != nil {
		t.Error(err)
	}

	// 测试用例
	DocTag := engine.DocTag{}
	json.Unmarshal(bytes, &DocTag)
	d := DocTag.Data.Item

	tagItem := model.TagItem{}
	tagItem.Biz = d.Biz
	tagItem.Category = d.Category
	tagItem.Type = d.Type

	var tags []model.Tags
	for _, v := range d.Tags {
		i := model.Tags{Category: v.Category, Type: v.Type, Name: v.Name, Text: v.Text, Tag: v.Tag}
		tags = append(tags, i)
	}

	tagItem.PosterUid = d.PosterUid
	tagItem.Tags = tags
	tagItem.UploadTime = d.UploadTime
	tagItem.UploadTimestamp = d.UploadTimestamp
	tagItem.Description = d.Description
	tagItem.AlreadyCollected = d.AlreadyCollected
	tagItem.AlreadyLiked = d.AlreadyLiked
	tagItem.UserStatus = d.UserStatus
	tagItem.ViewCount = d.ViewCount
	tagItem.LikeCount = d.LikeCount
	tagItem.CollectCount = d.CollectCount
	tagItem.VerifyStatus = d.VerifyStatus
	tagItem.AlreadyVoted = d.AlreadyVoted
	tagItem.VoteCount = d.VoteCount
	tagItem.CommentCount = d.CommentCount
	tagItem.CopyForbidden = d.Settings.CopyForbidden

	// 要测试的方法
	parseResult := h_bilibili_com.DocDetail(bytes, "2532912")

	for _, v := range parseResult.Items {

		if fmt.Sprint(tagItem) != fmt.Sprint(v.Payload) {
			t.Error("not match")
			fmt.Println(tagItem)
			fmt.Println("---------")
			fmt.Println(v.Payload)
		}

	}

}
