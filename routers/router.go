package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1 "github.com/xd-meal-back-end/api/v1"
	_ "github.com/xd-meal-back-end/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设置session midddleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/ImportUser", v1.ImportUser)
		apiv1.GET("/GetDishes", v1.GetDishes)
		apiv1.GET("/DownloadMenu", v1.DownloadMenu)
		apiv1.POST("/ReadMenu", v1.ReadMenu)
		apiv1.POST("/ImportMenu", v1.ImportMenu)
		apiv1.POST("/OrderDishes", v1.OrderDishes)
		apiv1.POST("/Login", v1.Login)
		apiv1.POST("/LoginOut", v1.LoginOut)
		apiv1.POST("/CheckUserLogin", v1.CheckUserLogin)
		apiv1.GET("/GetOrderDishes", v1.GetOrderDishes)
		apiv1.POST("/UpdateUserOrder", v1.UpdateUserOrder)
		apiv1.GET("/GetOrderSwitch", v1.GetOrderSwitch)
		apiv1.POST("/EnableOrderSwitch", v1.EnableOrderSwitch)
		apiv1.GET("/GetUserOrderSwitch", v1.GetUserOrderSwitch)
		apiv1.POST("/ResetPasswordByUser", v1.ResetPasswordByUser)
	}
	return r
}
