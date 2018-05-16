package fetch_test

import (
	"LearnGo/awe/spy/engine"
	"encoding/json"
	"testing"

	"LearnGo/awe/spy2/fetcher"
	"fmt"
	"io/ioutil"
	"strings"
)

func TestJsonPostFetch(t *testing.T) {

	info := `{
  "status": true,
  "data": {
    "mid": "25000899",
    "name": "狩子和光哥",
    "approve": false,
    "sex": "女",
    "rank": "10000",
    "face": "http://i0.hdslb.com/bfs/face/xxxxxxx.jpg",
    "DisplayRank": "10000",
    "regtime": 1456668728,
    "spacesta": 2,
    "birthday": "0000-00-00",
    "place": "广东省 广州市",
    "description": "",
    "article": 0,
    "sign": ".",
    "level_info": {
      "current_level": 5,
      "current_min": 10800,
      "current_exp": 17668,
      "next_exp": 28800
    },
    "pendant": {
      "pid": 0,
      "name": "",
      "image": "",
      "expire": 0
    },
    "nameplate": {
      "nid": 0,
      "name": "",
      "image": "",
      "image_small": "",
      "level": "",
      "condition": ""
    },
    "official_verify": {
      "type": -1,
      "desc": ""
    },
    "vip": {
      "vipType": 0,
      "vipDueDate": 0,
      "dueRemark": "",
      "accessStatus": 1,
      "vipStatus": 0,
      "vipStatusWarn": ""
    },
    "toutu": "bfs/space/xxxxxxxxx.png",
    "toutuId": 1,
    "theme": "default",
    "theme_preview": "",
    "coins": 0,
    "im9_sign": "",
    "playNum": 1259553,
    "fans_badge": true
  }
}`
	u := engine.UserInfo{}
	reader := strings.NewReader(info)
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(all, &u)
	fmt.Println(u)

	bytes, err := fetcher.JsonPostFetch("25000899")
	if err != nil {
		t.Error(err)
	}

	j := engine.UserInfo{}
	json.Unmarshal(bytes, &j)
	fmt.Println(j)

	if u.Status != j.Status {
		t.Errorf("status not match, need %v, %v get")
	}

	if u.Data.Mid != j.Data.Mid {
		t.Errorf("not match, need %v, %v get", u.Data.Mid, j.Data.Mid)
	}

	if u.Data.Name != j.Data.Name {
		t.Errorf("not match, need %v, %v get", u.Data.Name, j.Data.Name)
	}

	if u.Data.RegTime != j.Data.RegTime {
		t.Errorf("not match, need %v, %v get", u.Data.RegTime, j.Data.RegTime)
	}
}
