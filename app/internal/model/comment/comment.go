package comment

import (
	"juejin/app/internal/model/user"
	"time"
)

type Comment struct {
	CommentId      string       `json:"comment_id"`
	CommentContent string       `json:"comment_content"`
	CommentReplies []ReplyBrief `json:"comment_replies"`
	DiggCount      int          `json:"digg_count"`
	ReplyCount     int          `json:"reply_count"`
	IsDigg         bool         `json:"is_digg"`
	CreatTime      time.Time    `json:"creat_time"`
	ItemId         string       `json:"item_id"`
	ItemType       int          `json:"item_type"`
	UserId         int64        `json:"user_id"`
}

type ReplyBrief struct {
	ReplyId          string    `json:"reply_id"`
	ReplyToCommentId string    `json:"reply_to_comment_id"`
	ReplyToReplyId   string    `json:"reply_to_reply_id"`
	ReplyToUserId    string    `json:"reply_to_user_id"`
	ReplyContent     string    `json:"reply_content"`
	UserId           int64     `json:"user_id"`
	ItemId           string    `json:"item_id"`
	ItemType         int       `json:"item_type"`
	DiggCount        int       `json:"digg_count"`
	IsDigg           bool      `json:"is_digg"`
	CreatTime        time.Time `json:"creat_time"`
}

type ReplyInfo struct {
	ParentReplyInfo ReplyBrief    `json:"parent_reply_info"`
	ReplyInfo       ReplyBrief    `json:"reply_info"`
	ReplyToUserInfo user.InfoPack `json:"reply_to_user_info"`
	UserInfo        user.InfoPack `json:"user_info"`
}

type List struct {
	CommentInfo Comment       `json:"comment_info"`
	ReplyInfo   []ReplyInfo   `json:"reply_info"`
	UserInfo    user.InfoPack `json:"user_info"`
}
