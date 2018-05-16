package engine

// user lists
type UserListResult struct {
	Code    int        `json:"code"`
	Msg     string     `json:"msg"`
	Message string     `json:"message"`
	Data    DataResult `json:"data"`
}

type DataResult struct {
	Count int           `json:"total_count"`
	Items []ItemsResult `json:"items"`
}

type ItemsResult struct {
	User UserResult `json:"user"`
	Item ItemResult `json:"item"`
}

type UserResult struct {
	Uid     int    `json:"uid"`
	HeadUrl string `json:"head_url"`
	Name    string `json:"name"`
}

type ItemResult struct {
	DocId        int         `json:"doc_id"`
	PosterUid    int         `json:"poster_uid"`
	Pictures     interface{} `json:"pictures"`
	Title        string      `json:"title"`
	DocType      string      `json:"category"`
	UploadTime   int         `json:"upload_time"`
	AlreadyLiked int         `json:"already_liked"`
	AlreadyVoted int         `json:"already_voted"`
}

//GetInfo

type UserInfo struct {
	Status bool         `json:"status"`
	Data   UserInfoData `json:"data"`
}
type UserInfoData struct {
	Mid            int                `json:"mid"`
	Name           string             `json:"name"`
	Approve        bool               `json:"approve"`
	Sex            string             `json:"sex"`
	Rank           int                `json:"rank"`
	Face           string             `json:"face"`
	RegTime        int                `json:"regtime"`
	Place          string             `json:"place"`
	DisplayRank    string             `json:"DisplayRank"`
	Spacesta       int                `json:"spacesta"`
	Description    string             `json:"description"`
	Article        int                `json:"article"`
	Sign           string             `json:"sign"`
	LevelInfo      UserLevel          `json:"level_info"`
	Pendant        UserPendant        `json:"pendant"`
	Nameplate      UserNameplate      `json:"nameplate"`
	OfficialVerify UserOfficialVerify `json:"official_verify"`
	Vip            UserVip            `json:"vip"`
	Toutu          string             `json:"toutu"`
	ToutuId        int                `json:"toutu_id"`
	Theme          string             `json:"theme"`
	ThemePreview   string             `json:"theme_preview"`
	Coins          int                `json:"coins"`
	Im9Sign        string             `json:"im9_sign"`
	PlayNum        int                `json:"playNum"`
	FansBadge      bool               `json:"fans_badge"`
}

type UserLevel struct {
	CurrentLevel int `json:"current_level"`
	CurrentMin   int `json:"current_min"`
	CurrentExp   int `json:"current_exp"`
	NextExp      int `json:"next_exp"`
}

type UserPendant struct {
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Expire int    `json:"expire"`
}

type UserNameplate struct {
	Nid        int    `json:"nid"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	ImageSmall string `json:"image_small"`
	Level      string `json:"level"`
	Condition  string `json:"condition"`
}

type UserOfficialVerify struct {
	Type int    `json:"type"`
	Desc string `json:"desc"`
}

type UserVip struct {
	VipType       int `json:"vipType"`
	VipDueDate    int `json:"vipDueDate"`
	DueRemark     int `json:"dueRemark"`
	AccessStatus  int `json:"accessStatus"`
	VipStatus     int `json:"vipStatus"`
	VipStatusWarn int `json:"vipStatusWarn"`
}

// upload count
type UploadCount struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Message string          `json:"message"`
	Data    UploadCountData `json:"data"`
}
type UploadCountData struct {
	AllCount   int `json:"all_count"`
	DrawCount  int `json:"draw_count"`
	PhotoCount int `json:"photo_count"`
	DailyCount int `json:"daily_count"`
}

// doc_list
type DocList struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    DocListData `json:"data"`
}

type DocListData struct {
	Items []DocListDataItems `json:"items"`
}

type DocListDataItems struct {
	DocId       int        `json:"doc_id"`
	PosterUid   int        `json:"poster_uid"`
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
	Mid  string `json:"mid"`
	Csrf string `json:"csrf"`
}

//func (UserListResult) UnmarshalJSON(data []byte) error {
//	fmt.Println(string(data))
//	return nil
//}

type DocTag struct {
	Code    int        `json:"code"`
	Msg     string     `json:"msg"`
	Message string     `json:"message"`
	Data    DocTagItem `json:"data"`
}

type DocTagItem struct {
	Item TagItem `json:"item"`
}

type TagItem struct {
	PosterUid        int         `json:"poster_uid"`
	Biz              int         `json:"biz"`
	Category         string      `json:"category"`
	Type             int         `json:"type"`
	Tags             []Tags      `json:"tags"`
	UploadTime       string      `json:"upload_time"`
	UploadTimestamp  int         `json:"upload_timestamp"`
	Description      string      `json:"description"`
	AlreadyCollected int         `json:"already_collected"`
	AlreadyLiked     int         `json:"already_liked"`
	UserStatus       int         `json:"user_status"`
	ViewCount        int         `json:"view_count"`
	LikeCount        int         `json:"like_count"`
	CollectCount     int         `json:"collect_count"`
	VerifyStatus     int         `json:"verify_status"`
	AlreadyVoted     int         `json:"already_voted"`
	VoteCount        int         `json:"vote_count"`
	CommentCount     int         `json:"comment_count"`
	Settings         DocSettings `json:"settings"`
}

type Tags struct {
	Tag      string `json:"tag"`
	Type     int    `json:"type"`
	Category string `json:"category"`
	Text     string `json:"text"`
	Name     string `json:"name"`
}

type DocSettings struct {
	CopyForbidden int `json:"copy_forbidden"`
}
