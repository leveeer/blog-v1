package vo

type ArticleVO struct {
	Keyword    string `json:"keyword" form:"keyword"`
	Likes      string `json:"likes" form:"likes"`
	State      int  `json:"state" form:"state"`
	TagId      string `json:"tag_id" form:"tag_id"`
	CategoryId string `json:"category_id" form:"category_id"`
	PageNum    int `json:"pageNum" form:"pageNum"`
	PageSize   int `json:"pageSize" form:"pageSize"`
}
