package model

import "encoding/json"

type UserInfo struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Place     string `json:"place"`
	RegTime   int    `json:"reg_time"`
	Face      string `json:"face"`
	Rank      string `json:"rank"`
	Sign      string `json:"sign"`
	Level     int    `json:"level"`
	VipType   int    `json:"vip_type"`
	VipStatus int    `json:"vip_status"`
	Im9Sign   string `json:"im_9_sign"`
}

type UserUploadCount struct {
	DllCount   int `json:"dll_count"`
	DrawCount  int `json:"draw_count"`
	PhotoCount int `json:"photo_count"`
	DailyCount int `json:"daily_count"`
	User       int `json:"user"`
}

type UserDoc struct {
	Items []DocList `json:"items"`
}

type DocListLittle struct {
	DocId     string `json:"doc_id"`
	PosterUid string `json:"user"`
	Category  string `json:"cate"`
}

type DocList struct {
	DocId       int        `json:"doc_id"`
	PosterUid   int        `json:"user"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Pictures    []Pictures `json:"pictures"`
	Count       int        `json:"count"`
	Ctime       int        `json:"ctime"`
	View        int        `json:"view"`
	Like        int        `json:"like"`
}

type Pictures struct {
	ImgSrc    string `json:"img_src"`
	ImgWidth  int    `json:"img_width"`
	ImgHeight int    `json:"img_height"`
	ImgSize   int    `json:"img_size"`
}

// json
type PostBody struct {
	Mid  string
	Csrf string
}

func FromJsonObj(o interface{}) (UserInfo, error) {
	var userInfo UserInfo
	// 转换成string
	s, err := json.Marshal(o)
	if err != nil {
		return userInfo, err
	}
	// 转换成userinfo结构
	err = json.Unmarshal(s, &userInfo)
	return userInfo, err
}
