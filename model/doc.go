package model

type TagItem struct {
	PosterUid        int    `json:"poster_uid"`
	Biz              int    `json:"biz"`
	Category         string `json:"attr"`
	Type             int    `json:"type"`
	Tags             []Tags `json:"tags"`
	UploadTime       string `json:"upload_time"`
	UploadTimestamp  int    `json:"upload_timestamp"`
	Description      string `json:"description"`
	AlreadyCollected int    `json:"already_collected"`
	AlreadyLiked     int    `json:"already_liked"`
	UserStatus       int    `json:"user_status"`
	ViewCount        int    `json:"view_count"`
	LikeCount        int    `json:"like_count"`
	CollectCount     int    `json:"collect_count"`
	VerifyStatus     int    `json:"verify_status"`
	AlreadyVoted     int    `json:"already_voted"`
	VoteCount        int    `json:"vote_count"`
	CommentCount     int    `json:"comment_count"`
	CopyForbidden    int    `json:"copy_forbidden"`
}

type Tags struct {
	Tag      string `json:"tag"`
	Type     int    `json:"type"`
	Category string `json:"category"`
	Text     string `json:"text"`
	Name     string `json:"name"`
}
