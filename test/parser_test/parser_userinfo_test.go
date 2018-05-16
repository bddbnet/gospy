package parser_test

import (
	"LearnGo/awe/spy2/engine"
	"LearnGo/awe/spy2/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"LearnGo/awe/spy2/parser/h.bilibili.com"
)

// step 3 用户信息
func TestUserList(t *testing.T) {

	// 读取测试数据
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
	for _, v := range parseResult.Items {

		if userInfo != v.Payload {
			t.Error("not match")
			fmt.Println(userInfo)
			fmt.Println("---------")
			fmt.Println(v.Payload)
		}
	}

}

func TestPost(t *testing.T) {

	s := `
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="mid"

%s
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="csrf"

%s
------WebKitFormBoundary7MA4YWxkTrZu0gW--`
	payload := strings.NewReader(fmt.Sprintf(s, "8779497", "123456"))
	req, _ := http.NewRequest("POST", "http://space.bilibili.com/ajax/member/GetInfo", payload)

	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Set("Referer", fmt.Sprintf("https://space.bilibili.com/%s/", "3311852"))
	req.Header.Set("Host", "space.bilibili.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")
	req.Header.Set("Cache-Control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	userInfo := engine.UserInfo{}
	json.Unmarshal(body, &userInfo)

	fmt.Println(userInfo)

}
