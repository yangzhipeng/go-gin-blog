package routes

import (
	"gin-blog/internal/app/controller/v1"
	"gin-blog/internal/pkg/core"
	"gin-blog/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(HandleNotFound)
	router.POST("api/v1/user/register", v1.CreateUser) //用户注册
	router.POST("api/v1/user/login", v1.UserLogin)     //用户登陆

	r := router.Group("api/v1")
	r.Use(middleware.JwtAuth())          //使用jwt中间件
	r.POST("/user/:id", v1.UpdateUser)   //更新用户
	r.GET("/user", v1.UserList)          //用户列表
	r.POST("/user", v1.CreateUser)       //创建用户
	r.DELETE("/user/:id", v1.DeleteUser) //删除用户

	r.GET("/article", v1.ArticleList)          //文章列表
	r.POST("/article", v1.CreateArticle)       //创建文章
	r.POST("/article/:id", v1.UpdateArticle)   //更新文章
	r.DELETE("/article/:id", v1.DeleteArticle) //删除文章

	r.GET("/tag", v1.TagList)          //标签列表
	r.POST("/tag", v1.CreateTag)       //创建标签
	r.POST("/tag/:id", v1.UpdateTag)   //更新标签
	r.DELETE("/tag/:id", v1.DeleteTag) //删除标签

	r.GET("/category", v1.CategoryList)          //分类列表
	r.POST("/category", v1.CreateCategory)       //创建分类
	r.POST("/category/:id", v1.UpdateCategory)   //更新分类
	r.DELETE("/category/:id", v1.DeleteCategory) //删除分类

	return router
}

func HandleNotFound(c *gin.Context) {
	core.Response(nil, core.Code(404), core.Message("oops~ page not found")).Json(c)
	return
}
