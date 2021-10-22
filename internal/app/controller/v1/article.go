package v1

import (
	. "gin-blog/internal/pkg/common"
	"gin-blog/internal/pkg/core"
	"gin-blog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 文章列表
//@method GET
//@url /api/v1/article
//@param limit
//@param page
func ArticleList(c *gin.Context) {
	s := &service.ArticleService{}
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
type AddArticleParams struct {
	Title      string `json:"title" validate:"required"`              // 标题
	TagID      int    `json:"tag_id" validate:"required,min=0"`       // 标签id
	Status     int    `json:"status" validate:"required,min=0,max=1"` // 状态
	Desc       string `json:"desc"`                                   // 描述
	Content    string `json:"content" validate:"required"`            // 内容
	CategoryID int    `json:"category_id" validate:"required,min=0"`  // 分类id
}

// 创建文章
//@method POST
//@url /api/v1/article
//@param title
//@param tag_id
//@param status
//@param desc
//@param content
//@param category_id
func CreateArticle(c *gin.Context) {
	var params AddArticleParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	if tagIsExist := checkTag(c, params.TagID); !tagIsExist {
		return
	}
	if categoryIsExist := checkCategory(c, params.CategoryID); !categoryIsExist {
		return
	}
	s := &service.ArticleService{}
	s = addAssignValue(s, params)
	s.UserID = c.MustGet("uid").(int)
	if err := s.AddArticle(); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("insert failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

// 创建参数赋值
func addAssignValue(s *service.ArticleService, params AddArticleParams) *service.ArticleService {
	s.Title = params.Title
	s.TagID = params.TagID
	s.Status = params.Status
	s.Desc = params.Desc
	s.Content = params.Content
	s.CategoryID = params.CategoryID
	return s
}

// 更新表单参数
type UpdateArticleParams struct {
	Title      string `json:"title" validate:"required"`              // 标题
	TagID      int    `json:"tag_id" validate:"required,min=0"`       // 标签id
	Status     int    `json:"status" validate:"required,min=0,max=1"` // 状态
	Desc       string `json:"desc"`                                   // 描述
	Content    string `json:"content" validate:"required"`            // 内容
	CategoryID int    `json:"category_id" validate:"required,min=0"`  // 分类id
}

// 更新文章
//@method POST
//@url /api/v1/article/:id
//@param :id
//@param title
//@param tag_id
//@param status
//@param desc
//@param content
//@param category_id
func UpdateArticle(c *gin.Context) {
	var params UpdateArticleParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	if tagIsExist := checkTag(c, params.TagID); !tagIsExist {
		return
	}
	if categoryIsExist := checkCategory(c, params.CategoryID); !categoryIsExist {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	s := &service.ArticleService{}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("article not found")).Json(c)
		return
	}
	s = updateAssignValue(s, params)
	if err = s.EditArticle(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("update failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

// 更新参数赋值
func updateAssignValue(s *service.ArticleService, params UpdateArticleParams) *service.ArticleService {
	s.Title = params.Title
	s.TagID = params.TagID
	s.Status = params.Status
	s.Desc = params.Desc
	s.Content = params.Content
	s.CategoryID = params.CategoryID
	return s
}

// 检查标签是否存在
func checkTag(c *gin.Context, tid int) bool {
	if tid > 0 {
		if res, _ := (&service.TagService{}).FindById(tid); !res {
			core.Response("", core.Code(http.StatusBadRequest), core.Message("tag not found")).Json(c)
			return false
		}
	}
	return true
}

// 检查分类是否存在
func checkCategory(c *gin.Context, cid int) bool {
	if cid > 0 {
		if res, _ := (&service.CategoryService{}).FindById(cid); !res {
			core.Response("", core.Code(http.StatusBadRequest), core.Message("category not found")).Json(c)
			return false
		}
	}
	return true
}

// 删除文章
//@method DELETE
//@url /api/v1/article/:id
//@param :id
func DeleteArticle(c *gin.Context) {
	s := &service.ArticleService{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("article not found")).Json(c)
		return
	}
	if err = s.DelArticle(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("delete failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}
