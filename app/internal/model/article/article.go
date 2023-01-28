package article

import "time"

type Article struct {
	Content      string    `json:"content"`
	BriefContent string    `json:"brief_content"`
	Title        string    `json:"title"`
	CategoryId   string    `json:"category_id"`
	ArticleId    string    `json:"article_id"`
	UserId       string    `json:"user_id"`
	TagsIds      []string  `json:"tags_ids"`
	PublishTime  time.Time `json:"publish_time"`
}
