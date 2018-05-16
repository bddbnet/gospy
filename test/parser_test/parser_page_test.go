package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/bddbnet/gospy/parser/h.bilibili.com"
)

// step 2 获取用户列表
func TestCosPageList(t *testing.T) {

	// 用例
	const resultSize = 20

	expectDocId := []string{
		"3171340", "3174775", "3050585",
	}

	bytes, err := ioutil.ReadFile("cos_page.json")
	if err != nil {
		t.Error(err)
	}
	parseResult := h_bilibili_com.CosPageList(bytes)
	if len(parseResult.Items) != resultSize {
		t.Errorf("item need %d, %d give", resultSize, len(parseResult.Items))
	}

	for i, v := range expectDocId {
		if parseResult.Items[i].Id != v {
			t.Errorf("need %v, give %v\n", v, parseResult.Items[i].Id)
		}
	}

}
