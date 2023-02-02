package article

import (
	"juejin/app/internal/model/user"
	"time"
)

type Article struct {
	Content      string    `json:"content"`
	BriefContent string    `json:"brief_content"`
	Title        string    `json:"title"`
	CategoryId   string    `json:"category_id"`
	ArticleId    string    `json:"article_id"`
	UserId       string    `json:"user_id"`
	TagsIds      []string  `json:"tags_ids"`
	PublishTime  time.Time `json:"publish_time"`
	CollectCount int       `json:"collect_count"`
	CommentCount int       `json:"comment_count"`
	DiggCount    int       `json:"digg_count"`
	ViewCount    int       `json:"view_count"`
}

type Brief struct {
	ArticleId   string         `json:"article_id"`
	ArticleInfo *Article       `json:"article_info"`
	AuthorInfo  *user.InfoPack `json:"author_info"`
}
