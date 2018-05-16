package test

import (
	"testing"

	"context"

	"github.com/bddbnet/gospy/engine"
	"github.com/bddbnet/gospy/model"
	"github.com/bddbnet/gospy/parser/h.bilibili.com"
	"github.com/bddbnet/gospy/persist"

	"encoding/json"

	"fmt"
	"io/ioutil"

	"strconv"

	"gopkg.in/olivere/elastic.v5"
)

func TestItemSaver(t *testing.T) {
	excpected := engine.Item{
		Id:   "9904489",
		Url:  "",
		Type: "userinfo",
		Payload: model.UserInfo{
			Uid:       "9904489",
			Name:      "Aluckyysw",
			Sex:       "保密",
			Place:     "0000-01-01",
			RegTime:   1430153171,
			Face:      "http://i2.hdslb.com/bfs/face/3b073514c0c759d0ee3c0cbd785fc759823a0467.jpg",
			Rank:      "10000",
			Sign:      "",
			Level:     4,
			VipType:   0,
			VipStatus: 0,
			Im9Sign:   "47857cbe414ef07f31ed0ab1533f39b2",
		},
	}

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}

	const index = "t_1"
	err = persist.Save(client, index, excpected)

	if err != nil {
		t.Error(err)
	}

	resp, err := client.Get().Index(index).Type(excpected.Type).Id(excpected.Id).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual.Payload)
	if err != nil {
		t.Error(err)
	}

	//// map -> struct
	o, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = o

	if actual.Payload != excpected.Payload {
		t.Errorf("got  %v,\nneed %v\n", actual, excpected)
	}
}

func TestUserInfoSave(t *testing.T) {
	bytes, err := ioutil.ReadFile("userInfo.json")
	if err != nil {
		t.Error(err)
	}

	// 测试用例
	j := engine.UserInfo{}
	json.Unmarshal(bytes, &j)
	userInfo := model.UserInfo{}
	userInfo.Uid = strconv.Itoa(j.Data.Mid)
	userInfo.Name = j.Data.Name
	userInfo.Sex = j.Data.Sex
	userInfo.Place = j.Data.Place
	userInfo.RegTime = j.Data.RegTime
	userInfo.Rank = strconv.Itoa(j.Data.Rank)
	userInfo.Face = j.Data.Face
	userInfo.Sign = j.Data.Sign
	userInfo.Level = j.Data.LevelInfo.CurrentLevel
	userInfo.VipType = j.Data.Vip.VipType
	userInfo.VipStatus = j.Data.Vip.VipStatus
	userInfo.Im9Sign = j.Data.Im9Sign

	// 要测试的方法
	parseResult := h_bilibili_com.UserInfo(bytes, "88128280")

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parseResult.Items)
	const index = "t_1"
	err = persist.Save(client, index, engine.Item{Payload: parseResult.Items[0], Id: "25000899", Type: "userinfo"})
	if err != nil {
		panic(err)
	}
	for _, v := range parseResult.Items {

		if userInfo != v.Payload {
			t.Error("not match")
			fmt.Println(userInfo)
			fmt.Println("---------")
			fmt.Println(v.Payload)
		}
	}

}

func TestDocSave(t *testing.T) {
	// 测试数据
	bytes, err := ioutil.ReadFile("doc_list.json")
	if err != nil {
		t.Error(err)
	}

	// 测试用例
	j := engine.DocList{}
	json.Unmarshal(bytes, &j)

	// 测试方法
	parseResult := h_bilibili_com.DocLists(bytes)

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parseResult.Items)
	const index = "t_1"
	for _, item := range parseResult.Items {
		err = persist.Save(client, index, engine.Item{Payload: item, Id: "2586430", Type: "userdoc"})
		if err != nil {
			panic(err)
		}
	}

	for _, item := range parseResult.Items {

		fmt.Println(item)
	}
}
