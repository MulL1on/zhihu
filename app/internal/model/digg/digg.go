package digg

type Like struct {
	AuthorId string `json:"author_id"`
	ItemId   string `json:"item_id"`
	ItemType int    `json:"item_type"`
}
