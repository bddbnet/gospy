package parser_test

import (
	"LearnGo/awe/spy2/engine"
	"LearnGo/awe/spy2/model"
	"encoding/json"
	"io/ioutil"
	"testing"

	"fmt"

	"LearnGo/awe/spy2/parser/h.bilibili.com"
)

// step 5 图片列表
func TestDocLists(t *testing.T) {
	// 测试数据
	bytes, err := ioutil.ReadFile("doc_list.json")
	if err != nil {
		t.Error(err)
	}

	// 测试用例
	j := engine.DocList{}
	json.Unmarshal(bytes, &j)
	//fmt.Println(j.Data.Items)

	// 测试方法
	parseResult := h_bilibili_com.DocLists(bytes)

	if len(j.Data.Items) != len(parseResult.Items) {
		t.Error("not match")
	}

	for k, v := range j.Data.Items {
		d := parseResult.Items[k].Payload.(model.DocList)
		if d.Title != v.Title {
			t.Errorf("need %s, %s get", d.Title, v.Title)
		}
		if d.Ctime != v.Ctime {
			t.Errorf("need %s, %s get", d.Ctime, v.Ctime)
		}
		fmt.Println(d)
		fmt.Println(v)
		if fmt.Sprint(d) != fmt.Sprint(v) {
			t.Errorf("need \n%s, \n%s get\n", fmt.Sprint(d), fmt.Sprint(v))
		}

	}

}
