package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1 "github.com/xd-meal-back-end/api/v1"
	_ "github.com/xd-meal-back-end/docs"
	"github.com/xd-meal-back-end/middleware/auth"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设置session midddleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//登录接口
	r.POST("/api/v1/Login", v1.Login)
	r.POST("/api/v1/LoginOut", v1.LoginOut)
	r.GET("/api/v1/GetQRCode", v1.GetQRCode)
	r.GET("/api/v1/WeiXinLogin", v1.WeiXinLogin)
	r.GET("/api/v1/CheckUserLogin", v1.CheckUserLogin)
	//pos机接口
	r.POST("/api/v1/ScanDishCode", v1.ScanDishCode)
	//用户接口
	apiv1 := r.Group("/api/v1")
	apiv1.Use(auth.CheckAuth())
	{
		apiv1.GET("/GetDishes", v1.GetDishes)
		apiv1.POST("/OrderDishes", v1.OrderDishes)
		apiv1.GET("/GetOrderDishes", v1.GetOrderDishes)
		apiv1.POST("/UpdateUserOrder", v1.UpdateUserOrder)
		apiv1.GET("/GetUserOrderSwitch", v1.GetUserOrderSwitch)
		apiv1.POST("/ResetPasswordByUser", v1.ResetPasswordByUser)
		apiv1.GET("/GetDishCode", v1.GetDishCode)
		apiv1.POST("/EvalDish", v1.EvalDish)
	}
	//后台接口
	apiadmin := r.Group("/api/v1")
	apiadmin.Use(auth.CheckAdminAuth())
	{
		apiadmin.POST("/ImportUser", v1.ImportUser)               //导入外部用户
		apiadmin.GET("/DownloadMenu", v1.DownloadMenu)            //下载菜单模版
		apiadmin.POST("/ImportMenu", v1.ImportMenu)               //导入下周菜单
		apiadmin.POST("/ReadMenu", v1.ReadMenu)                   //预览导入
		apiadmin.GET("/GetUserByEmail", v1.GetUserByEmail)        //根据邮箱获取用户信息
		apiadmin.POST("/AddMenuSingle", v1.AddMenuSingle)         //单独加菜
		apiadmin.GET("/GetMealTotal", v1.GetMealTotal)            //根据菜品统计订餐数量
		apiadmin.POST("/EnableOrderSwitch", v1.EnableOrderSwitch) //开启&关闭订餐开关
		apiadmin.GET("/GetOrderSwitch", v1.GetOrderSwitch)        //获取订餐开关状态
		apiadmin.POST("/ResetPassword", v1.ResetPassword)         //后台根据id重置密码
	}
	return r
}
