package vo

import (
	"blog-go-gin/models"
)

type BlogVO struct {
	Keyword           string          `form:"keyword" json:"keyword"`
	CurrentPage       int             `form:"currentPage" json:"currentPage"`
	PageSize          int             `form:"pageSize" json:"pageSize"`
	Uid               string          `form:"uid" json:"uid"`
	Status            int             `form:"status" json:"status"`
	Title             string          `form:"title" json:"title"`
	Summary           string          `form:"summary" json:"summary"`
	TagUid            string          `form:"tagUid" json:"tagUid"`
	BlogSortUid       string          `form:"blogSortUid" json:"blogSortUid"`
	FileUid           string          `form:"fileUid" json:"fileUid"`
	AdminUid          string          `form:"adminUid" json:"adminUid"`
	IsPublish         string          `form:"isPublish" json:"isPublish"`
	IsOriginal        string          `form:"isOriginal" json:"isOriginal"`
	Author            string          `form:"author" json:"author"`
	ArticlesPart      string          `form:"articlesPart" json:"articlesPart"`
	Level             int             `form:"level" json:"level"`
	Type              interface{}     `form:"type" json:"type"`
	OutsideLink       string          `form:"outsideLink" json:"outsideLink"`
	Content           string          `form:"content" json:"content"`
	TagList           []models.Tag    `form:"tagList" json:"tagList"`
	PhotoList         []string        `form:"photoList" json:"photoList"`
	BlogSort          models.BlogSort `form:"blogSort" json:"blogSort"`
	ParseCount        int             `form:"parseCount" json:"parseCount"`
	Copyright         string          `form:"copyright" json:"copyright"`
	LevelKeyword      interface{}     `form:"levelKeyword" json:"levelKeyword"`
	UserSort          int             `form:"userSort" json:"userSort"`
	Sort              int             `form:"sort" json:"sort"`
	OpenComment       interface{}     `form:"openComment" json:"openComment"`
	OrderByDescColumn string          `form:"orderByDescColumn" json:"orderByDescColumn"`
	OrderByAscColumn  string          `form:"orderByAscColumn" json:"orderByAscColumn"`
}
