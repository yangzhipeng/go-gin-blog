package middleware

import (
	"gin-blog/internal/pkg/core"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtSecret = []byte("cms-plan-blog_task")

type Claims struct {
	Username string `json:"username"`
	Uid      int    `json:"uid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(username string, uid int, role string) (string, int) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	setClaim := Claims{
		Username: username,
		Uid:      uid,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaim) // 生成token
	token, err := reqClaim.SignedString(JwtSecret)                  // 转换为字符串
	if err != nil {
		return "", http.StatusBadRequest
	}
	return token, http.StatusOK
}

// 验证token
func CheckToken(token string) (*Claims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if key, _ := setToken.Claims.(*Claims); setToken.Valid {
		return key, http.StatusOK
	} else {
		return nil, http.StatusBadRequest
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			core.Response("token不存在", core.Code(http.StatusForbidden), core.Message("auth failed")).Json(c)
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			core.Response("token格式错误", core.Code(http.StatusForbidden), core.Message("auth failed")).Json(c)
			c.Abort()
			return
		}
		key, code := CheckToken(checkToken[1])
		if code == http.StatusBadRequest {
			core.Response("token错误", core.Code(http.StatusForbidden), core.Message("auth failed")).Json(c)
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			core.Response("token已过期", core.Code(http.StatusForbidden), core.Message("auth failed")).Json(c)
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Set("uid", key.Uid)
		c.Set("role", key.Role)
		c.Next()
	}
}
