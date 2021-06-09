package vo

import "blog-go-gin/models/model"

type ArticleVo struct {
	Article              *model.Article   `json:"article"`      //当前文章
	LastArticle          *model.Article   `json:"last_article"` //上一篇
	NextArticle          *model.Article   `json:"next_article"` //下一篇
	RecommendArticleList []*model.Article `json:"recommend_article_list"`
}
