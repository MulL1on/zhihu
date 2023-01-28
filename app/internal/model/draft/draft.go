package draft

import "time"

type Draft struct {
	Content      string    `json:"content"`
	BriefContent string    `json:"brief_content"`
	Title        string    `json:"title"`
	CategoryId   string    `json:"category_id"`
	DraftId      string    `json:"draft_id"`
	UserId       string    `json:"user_id"`
	TagsIds      []string  `json:"tags_ids"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}
