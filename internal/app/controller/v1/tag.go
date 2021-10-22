package v1

import (
	. "gin-blog/internal/pkg/common"
	"gin-blog/internal/pkg/core"
	"gin-blog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//标签列表
//@method GET
//@url /api/v1/tag
//@param limit
//@param page
func TagList(c *gin.Context) {
	s := &service.TagService{}
	err := c.ShouldBindQuery(s)
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("query params value is invalid")).Json(c)
		return
	}
	where := s.QueryParams()
	total, err := s.Count(where)
	if err != nil {
		core.Response(err, core.Code(http.StatusInternalServerError), core.Message("query error")).Json(c)
		return
	}
	pager := PageInfo(c, total)
	s.Limit = pager.Limit
	s.Page = pager.Page

	articles, err := s.FindList(where)
	if err != nil {
		core.Response(err, core.Code(http.StatusInternalServerError), core.Message("query error")).Json(c)
		return
	}

	pager.Rows = len(articles)
	core.Response(formatData(articles, pager), core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

// 创建表单参数
type AddTagParams struct {
	Name   string `json:"name" validate:"required"`               // 名称
	Status int    `json:"status" validate:"required,min=0,max=1"` // 状态
}

//创建标签
//@method POST
//@url /api/v1/tag
//@param name
//@param status
func CreateTag(c *gin.Context) {
	var params AddTagParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	s := &service.TagService{}
	s.Name = params.Name
	s.Status = params.Status
	if err := s.AddTag(); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("insert failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

// 更新表单参数
type UpdateTagParams struct {
	Name   string `json:"name" validate:"required"`               // 名称
	Status int    `json:"status" validate:"required,min=0,max=1"` // 状态
}

//更新标签
//@method POST
//@url /api/v1/tag/:id
//@param :id
//@param name
//@param status
func UpdateTag(c *gin.Context) {
	var params UpdateTagParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	s := &service.TagService{}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("tag not found")).Json(c)
		return
	}
	s.Name = params.Name
	s.Status = params.Status
	if err = s.EditTag(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("update failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

//删除标签
//@method DELETE
//@url /api/v1/tag/:id
//@param :id
func DeleteTag(c *gin.Context) {
	s := &service.TagService{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("tag not found")).Json(c)
		return
	}
	if err = s.DelTag(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("delete failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}
