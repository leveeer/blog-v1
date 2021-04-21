

package page


type IPage struct {
	Records          interface{} `json:"records"`
	Total            int64       `json:"total"`
	Size             int         `json:"size"`
	Current          int         `json:"current"`
	OptimizeCountSql bool        `json:"optimizeCountSql"`
	IsSearchCount    bool        `json:"isSearchCount"`
}
