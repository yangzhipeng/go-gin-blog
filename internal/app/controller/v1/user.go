package v1

import (
	"gin-blog/internal/models"
	. "gin-blog/internal/pkg/common"
	"gin-blog/internal/pkg/core"
	"gin-blog/internal/pkg/middleware"
	"gin-blog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//用户列表
//@method GET
//@url /api/v1/user
//@param limit
//@param page
func UserList(c *gin.Context) {
	s := &service.UserService{}
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
	core.Response(formatData(articles, pager)).Json(c)
}

// 创建表单参数
type AddUserParams struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	Desc     string `json:"desc"`                         // 描述
}

//创建用户、注册用户
//@method POST
//@url /api/v1/user/register
//@param username
//@param password
//@param desc
func CreateUser(c *gin.Context) {
	var params AddUserParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	s := &service.UserService{}
	if res, _ := s.FindByUserName(params.Username); res {
		core.Response("", core.Code(http.StatusBadRequest), core.Message("user already exist")).Json(c)
		return
	}
	s.Username = params.Username
	s.Password = params.Password
	s.Desc = params.Desc
	if err := s.AddUser(); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("register failed, try it again")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message("success")).Json(c)
}

// 更新表单参数
type UpdateUserParams struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	Desc     string `json:"desc"`                         // 描述
}

//更新用户
//@method POST
//@url /api/v1/user/:id
//@param :id
//@param username
//@param password
//@param desc
func UpdateUser(c *gin.Context) {
	var params UpdateUserParams
	_ = c.ShouldBindJSON(&params)
	if boolV := ValidateParams(c, params); !boolV {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	s := &service.UserService{}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("user not found")).Json(c)
		return
	}
	s.Username = params.Username
	s.Password = params.Password
	s.Desc = params.Desc
	if err = s.EditUser(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("update failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

//删除用户
//@method DELETE
//@url /api/v1/user/:id
//@param :id
func DeleteUser(c *gin.Context) {
	s := &service.UserService{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("id is invalid")).Json(c)
		return
	}
	if res, _ := s.FindById(id); !res {
		core.Response(err, core.Code(http.StatusNotFound), core.Message("user not found")).Json(c)
		return
	}
	if err = s.DelUser(id); err != nil {
		core.Response(err, core.Code(http.StatusBadRequest), core.Message("delete failed")).Json(c)
		return
	}

	core.Response("", core.Code(http.StatusOK), core.Message(http.StatusText(http.StatusOK))).Json(c)
}

//用户登陆
//@method POST
//@url /api/v1/user/login
//@param username
//@param password
func UserLogin(c *gin.Context) {
	s := &service.UserService{}
	var entity models.User
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		core.Response(err, core.Code(http.StatusUnprocessableEntity), core.Message("params value is invalid")).Json(c)
		return
	}
	code, user := s.CheckLogin(entity.Username, entity.Password)
	if code == http.StatusOK {
		token, co := middleware.GenerateToken(user.Username, user.ID, user.Role)
		if co == http.StatusOK {
			core.Response(gin.H{"token": token}, core.Code(code), core.Message("login success")).Json(c)
		} else {
			core.Response(gin.H{"token": token}, core.Code(code), core.Message("login failed")).Json(c)
		}
	}
	if code == http.StatusNotFound {
		core.Response("", core.Code(code), core.Message("user not found")).Json(c)
	}
	if code == http.StatusForbidden {
		core.Response("", core.Code(code), core.Message("password error")).Json(c)
	}
}
