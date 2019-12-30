package auth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xd-meal-back-end/pkg/e"
	"net/http"
)

type UserAuth struct {
}

/**
app用户
*/
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		logier := UserAuth{}.IsLogin(c)
		if logier == nil {
			code = e.ERROR_AUTH
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

/**
后台用户
*/
func CheckAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		logier, roleType := UserAuth{}.IsAdminLogin(c)
		fmt.Println(roleType)
		if logier == nil || roleType != 1 {
			code = e.ERROR_ADMIN_AUTH
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func (ud UserAuth) IsLogin(c *gin.Context) interface{} {
	session := sessions.Default(c)
	logier := session.Get("logier")
	if logier != nil {
		return logier
	} else {
		return nil
	}
}

func (ud UserAuth) IsAdminLogin(c *gin.Context) (interface{}, int32) {
	session := sessions.Default(c)
	logier := session.Get("logier")
	roleType := session.Get("roleType").(int32)
	if logier != "" {
		return logier, roleType
	} else {
		return "", 0
	}
}
