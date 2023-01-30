package collection

import "time"

type Set struct {
	CollectionId     int64     `json:"collection_id"`
	CollectionName   string    `json:"collection_name"`
	Description      string    `json:"description"`
	UserId           int64     `json:"user_id"`
	Permission       int       `json:"permission"`
	PostArticleCount int       `json:"post_article_count"`
	CreateTime       time.Time `json:"create_time"`
	UpdateTime       time.Time `json:"update_time"`
}

type SelectArticle struct {
	ArticleId       int   `json:"article_id"`
	CollectionSetId int64 `json:"collection_set_id"`
}
