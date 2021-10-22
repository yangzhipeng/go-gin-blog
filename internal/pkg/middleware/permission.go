package middleware

import (
	"gin-blog/internal/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)
		if role != "admin" {
			core.Response("权限不足", core.Code(http.StatusForbidden), core.Message("forbidden!")).Json(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
