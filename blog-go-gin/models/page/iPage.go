package page

import "gorm.io/gorm"

type IPage struct {
	Total            int64       `json:"total"`
	Size             int         `json:"size"`
	Current          int         `json:"current"`
	OptimizeCountSql bool        `json:"optimizeCountSql"`
	IsSearchCount    bool        `json:"isSearchCount"`
	Records          interface{} `json:"records"`
}

//分页封装
func Paginate(page *IPage) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Current == 0 {
			page.Current = 1
		}
		switch {
		case page.Size > 100:
			page.Size = 100
		case page.Size <= 0:
			page.Size = 10
		}
		offset := (page.Current - 1) * page.Size
		return db.Offset(offset).Limit(page.Size)
	}
}
