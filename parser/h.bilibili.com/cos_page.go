package h_bilibili_com

import (
	"LearnGo/awe/spy2/engine"
	"LearnGo/awe/spy2/fetcher"
	"LearnGo/awe/spy2/model"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"LearnGo/awe/spy2/config"

	"github.com/PuerkitoBio/goquery"
)

const TypeDoc = config.ElasticSearchTypeDoc
const TypeUser = config.ElasticSearchTypeUser
const TypeDocCount = config.ElasticSearchTypeDocCount

const cssSelect = `div#main>a`

// step 1 获取列表
func CosIndexList(byteData []byte) engine.ParseResult {

	reader := bytes.NewReader(byteData)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)

	}

	result := engine.ParseResult{}

	doc.Find(cssSelect).Each(func(i int, selection *goquery.Selection) {
		//a.app_impression_tracked:nth-child(1) > div:nth-child(3) > div:nth-child(1)

		s := selection.Text()
		url, _ := selection.Attr("href")

		item := engine.Item{Id: "", Url: url, Type: "", Payload: s}
		result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{Url: url, ParseFunc: CosPageList, FetchFunc: fetcher.JsonFetch})
	})
	return result
}

// parseFunc
// step 2 获取用户列表
func CosPageList(byteData []byte) engine.ParseResult {
	j := engine.UserListResult{}
	json.Unmarshal(byteData, &j)

	pr := engine.ParseResult{}
	for _, v := range j.Data.Items {
		uid := strconv.Itoa(v.User.Uid)
		docId := strconv.Itoa(v.Item.DocId)
		url := "https://space.bilibili.com/ajax/member/GetInfo/" + uid

		category := v.Item.DocType
		docListLittle := model.DocListLittle{PosterUid: uid, DocId: docId, Category: category}

		item := engine.Item{}
		item.Id = docId
		item.Url = url
		item.Type = "user"
		item.Payload = docListLittle
		item.Action = "index"

		pr.Items = append(pr.Items, item)

		pr.Requests = append(pr.Requests, engine.Request{Url: url, FetchFunc: func(s string) ([]byte, error) {
			return fetcher.JsonPostFetch(uid)
		}, ParseFunc: func(bytes []byte) engine.ParseResult {
			return UserInfo(bytes, uid)
		}})

	}

	return pr
}

// step 3 处理用户信息
func UserInfo(byteData []byte, uid string) engine.ParseResult {
	j := engine.UserInfo{}
	json.Unmarshal(byteData, &j)

	if j.Status == false {
		var pr engine.ParseResult
		pr.Requests = append(pr.Requests, engine.Request{Url: "", FetchFunc: fetcher.NilFetch, ParseFunc: engine.NilParser})
		return pr
	}

	userInfo := model.UserInfo{}
	userInfo.Uid = strconv.Itoa(j.Data.Mid)
	userInfo.Name = j.Data.Name
	userInfo.Sex = j.Data.Sex
	userInfo.RegTime = j.Data.RegTime
	userInfo.Place = j.Data.Place
	userInfo.Rank = strconv.Itoa(j.Data.Rank)
	userInfo.Face = j.Data.Face
	userInfo.Sign = j.Data.Sign
	userInfo.Level = j.Data.LevelInfo.CurrentLevel
	userInfo.VipType = j.Data.Vip.VipType
	userInfo.VipStatus = j.Data.Vip.VipStatus
	userInfo.Im9Sign = j.Data.Im9Sign

	url := fmt.Sprintf("https://api.vc.bilibili.com/link_draw/v1/doc/upload_count?uid=%s", string(uid))

	pr := engine.ParseResult{}
	item := engine.Item{Id: userInfo.Uid, Url: url, Type: TypeUser, Payload: userInfo, Action: "create"}

	pr.Items = append(pr.Items, item)
	pr.Requests = append(pr.Requests, engine.Request{Url: url, FetchFunc: fetcher.JsonFetch, ParseFunc: func(bytes []byte) engine.ParseResult {
		return UserUploadCount(bytes, uid)
	}})

	return pr
}

// step 4 获取用户图片总数
func UserUploadCount(byteData []byte, uid string) engine.ParseResult {
	j := engine.UploadCount{}
	json.Unmarshal(byteData, &j)

	pr := engine.ParseResult{}

	userUploadCount := model.UserUploadCount{}
	user, _ := strconv.Atoi(uid)
	userUploadCount.User = user
	userUploadCount.DllCount = j.Data.AllCount
	userUploadCount.DailyCount = j.Data.DailyCount
	userUploadCount.DrawCount = j.Data.DrawCount
	userUploadCount.PhotoCount = j.Data.PhotoCount

	pr.Items = []engine.Item{
		{
			Id:       uid,
			Url:      "xxx",
			Type:     TypeDocCount,
			Payload:  userUploadCount,
			Action:   "create",
			ParentId: uid,
		},
	}

	// 每页个数
	perPageLimit := 30
	if j.Data.AllCount > 0 {
		drawPage := Page(perPageLimit, j.Data.AllCount)
		for x := 0; x <= drawPage; x++ {
			drawUrl := fmt.Sprintf(`https://api.vc.bilibili.com/link_draw/v1/doc/doc_list?uid=%s&page_num=%d&page_size=%d&biz=all`, uid, x, perPageLimit)
			pr.Requests = append(pr.Requests, engine.Request{Url: drawUrl, FetchFunc: fetcher.JsonFetch, ParseFunc: DocLists})
		}
	}

	return pr
}

// step 5 图片列表
func DocLists(byteData []byte) engine.ParseResult {

	j := engine.DocList{}
	json.Unmarshal(byteData, &j)
	//fmt.Println(j)
	pr := engine.ParseResult{}

	for _, v := range j.Data.Items {
		var Pictures []model.Pictures
		for _, pic := range v.Pictures {
			p := model.Pictures{}
			p.ImgSize = pic.ImgSize
			p.ImgHeight = pic.ImgHeight
			p.ImgWidth = pic.ImgWidth
			p.ImgSrc = pic.ImgSrc
			Pictures = append(Pictures, p)
		}

		doc := model.DocList{
			DocId:       v.DocId,
			PosterUid:   v.PosterUid,
			Title:       v.Title,
			Description: v.Description,
			Pictures:    Pictures,
			Count:       v.Count,
			Ctime:       v.Ctime,
			View:        v.View,
			Like:        v.Like,
		}

		item := engine.Item{}
		item.Id = strconv.Itoa(v.DocId)
		item.Url = ""
		item.Type = TypeDoc

		item.Payload = doc
		item.Action = "create"
		item.ParentId = strconv.Itoa(v.PosterUid)

		url := fmt.Sprintf("https://api.vc.bilibili.com/link_draw/v1/doc/detail?doc_id=%s", item.Id)

		pr.Items = append(pr.Items, item)
		pr.Requests = append(pr.Requests, engine.Request{Url: url, FetchFunc: fetcher.JsonFetch, ParseFunc: func(bytes []byte) engine.ParseResult {
			return DocDetail(bytes, item.Id)
		}})

	}

	return pr
}

// step 6 doc 详情
func DocDetail(byteData []byte, docId string) engine.ParseResult {
	//https://api.vc.bilibili.com/link_draw/v1/doc/detail?doc_id=2532912
	docList := engine.DocTag{}
	json.Unmarshal(byteData, &docList)
	if docList.Code != 0 {
		// 未获取到结果 丢弃
		pr := engine.ParseResult{}
		pr.Requests = append(pr.Requests, engine.Request{Url: "", FetchFunc: fetcher.NilFetch, ParseFunc: engine.NilParser})
		return pr
	}
	d := docList.Data.Item

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

	pr := engine.ParseResult{}
	pr.Items = []engine.Item{
		{
			Id:       docId,
			Url:      "",
			Type:     TypeDoc,
			Action:   "update",
			Payload:  tagItem,
			ParentId: strconv.Itoa(d.PosterUid),
		},
	}
	pr.Requests = append(pr.Requests, engine.Request{Url: "", FetchFunc: fetcher.NilFetch, ParseFunc: engine.NilParser})
	return pr
}

// 分页
func Page(limit, total int) int {
	return (total+limit-1)/limit - 1
}
