package v1

import (
	. "gin-blog/internal/pkg/common"
	"gin-blog/internal/pkg/core"
	"gin-blog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//分类列表
//@method GET
//@url /api/v1/category
//@param limit
//@param page
func CategoryList(c *gin.Context) {
	s := &service.CategoryService{}
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
type AddCategoryParams struct {
	Name   string `json:"name" validate:"required"`               // 名称
	Status int    `json:"status" validate:"required,min=0,max=1"` // 状态
	Desc   string `json:"desc"`                                   // 描述
}

//创建分类
//@method POST
//@url /api/v1/category
//@param name
//@param status
//@param desc
func CreateCategory(c *gin.Context) {
	var params AddCategoryParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	s := &service.CategoryService{}
	s.Name = params.Name
	s.Status = params.Status
	s.Desc = params.Desc
	if err := s.AddCategory(); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("insert failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

// 更新表单参数
type UpdateCategoryParams struct {
	Name   string `json:"name" validate:"required"`               // 名称
	Status int    `json:"status" validate:"required,min=0,max=1"` // 状态
	Desc   string `json:"desc"`                                   // 描述
}

//更新分类
//@method POST
//@url /api/v1/category/:id
//@param :id
//@param name
//@param status
//@param desc
func UpdateCategory(c *gin.Context) {
	var params UpdateCategoryParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	s := &service.CategoryService{}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("category not found")).Json(c)
		return
	}
	s.Name = params.Name
	s.Status = params.Status
	s.Desc = params.Desc
	if err = s.EditCategory(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("update failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

//删除分类
//@method DELETE
//@url /api/v1/category/:id
//@param :id
func DeleteCategory(c *gin.Context) {
	s := &service.CategoryService{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("category not found")).Json(c)
		return
	}
	if err = s.DelCategory(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("delete failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}
