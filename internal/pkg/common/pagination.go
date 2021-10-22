package common

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

var (
	defaultLimit, defaultPage = 10, 1
)

type Pagination struct {
	Limit      int `form:"limit" json:"limit"`
	Page       int `form:"page" json:"page"`
	TotalPages int `json:"total_pages"`
	TotalRows  int `json:"total_rows"`
	Rows       int `json:"rows"`
}

// 分页信息
func PageInfo(c *gin.Context, total int) Pagination {
	var err error
	pager := Pagination{}
	pager.Limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil || pager.Limit <= 0 {
		pager.Limit = defaultLimit
	}
	pager.Page, err = strconv.Atoi(c.Query("page"))
	if err != nil || pager.Page <= 0 {
		pager.Page = defaultPage
	}
	pager.TotalRows = total
	pager.TotalPages = int(math.Ceil(float64(total) / float64(pager.Limit)))

	return pager
}

// 计算偏移量
func GetOffset(limit, page int) int {
	skip := 0
	if page > 0 {
		skip = (page - 1) * limit
	}

	return skip
}
