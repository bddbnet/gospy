package parser_test

import (
	"io/ioutil"
	"testing"

	"LearnGo/awe/spy2/parser/h.bilibili.com"
)

// step 1 获取列表
func TestIndexList(t *testing.T) {
	const resultSize = 24
	expectedItem := []string{
		"Page:  1", "Page:  2", "Page:  3",
	}
	expectedUrl := []string{
		"https://api.vc.bilibili.com/link_draw/v2/Photo/list?category=cos&type=hot&page_num=1&page_size=20",
		"https://api.vc.bilibili.com/link_draw/v2/Photo/list?category=cos&type=hot&page_num=2&page_size=20",
		"https://api.vc.bilibili.com/link_draw/v2/Photo/list?category=cos&type=hot&page_num=3&page_size=20",
	}

	file, err := ioutil.ReadFile("cos_index.html")
	if err != nil {
		t.Error(err)
	}
	parseResult := h_bilibili_com.CosIndexList(file)

	if len(parseResult.Items) != resultSize {
		t.Errorf("Item must have %d, %d found", resultSize, len(parseResult.Items))
	}
	if len(parseResult.Requests) != resultSize {
		t.Errorf("Re must have %d, %d found", resultSize, len(parseResult.Requests))
	}

	for i, item := range expectedItem {
		if parseResult.Items[i].Payload != item {
			t.Errorf("item need %s,%s found", item, parseResult.Items[i])
		}
	}
	for i, url := range expectedUrl {
		if parseResult.Requests[i].Url != url {
			t.Errorf("url need %s,%s found", url, parseResult.Requests[i].Url)
		}
	}
}
