package v1

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/arrays"
	"github.com/xd-meal-back-end/Function"
	"github.com/xd-meal-back-end/middleware"
	"github.com/xd-meal-back-end/middleware/mongo"
	"github.com/xd-meal-back-end/pkg/e"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

/**
导入外部用户
*/
func ImportUser(c *gin.Context) {
	//登录验证
	logier, roleType := UserData{}.isAdminLogin(c)
	if logier == nil || roleType != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "没有权限", "data": "",
		})
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	data, err := middleware.ImportUser(file)
	for _, v := range data {
		insert := mongo.UserMongo{ID: primitive.NewObjectID(), Name: v.Name, Email: v.Email,
			PassWord: fmt.Sprintf("%x", md5.Sum([]byte(v.PassWord))), Type: v.Type, Depart: v.Depart, CreateTime: v.CreateTime, Unique: v.Email}
		insert.CreateRow()
	}
	c.JSON(http.StatusOK, gin.H{"msg": err, "data": data})
}

/**
预览菜单
*/
func ReadMenu(c *gin.Context) {
	//登录验证
	logier, roleType := UserData{}.isAdminLogin(c)
	if logier == nil || roleType != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "没有权限", "data": "",
		})
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件错误"})
		return
	}
	//选饭启动区间区间
	filter := bson.M{"name": "order", "enable": 0}
	switches := mongo.Switches{}.FindOne(filter)
	if switches == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 401, "msg": "关闭选餐后才能导入菜单"})
		return
	}
	data, err := middleware.ReadMenuExcel(file)
	c.JSON(http.StatusOK, gin.H{"msg": err, "data": data})
}

/**
导入菜单
*/
func ImportMenu(c *gin.Context) {
	//登录验证
	logier, roleType := UserData{}.isAdminLogin(c)
	if logier == nil || roleType != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "没有权限", "data": "",
		})
		return
	}
	var param map[string][]mongo.DishesMongo
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  "参数错误",
			"data": "",
		})
		return
	}
	if param["data"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  "空数据",
			"data": "",
		})
		return
	}
	//选饭启动区间区间
	filter := bson.M{"name": "order", "enable": 0}
	switches := mongo.Switches{}.FindOne(filter)
	if switches == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 401, "msg": "关闭选餐后才能导入菜单"})
		return
	}
	var timeInterval []string
	for _, v := range param["data"] {
		timeInterval = append(timeInterval, v.MealDay)
		v.CreateRow()
		//合并菜品库
		filterDishLib := bson.M{"name": v.Name}
		exist := mongo.DishLib{}.FindOne(filterDishLib)
		if exist == nil {
			mongo.DishLib{ID: primitive.NewObjectID(), Name: v.Name, Supplier: v.Supplier, CreateTime: time.Now()}.CreateRow()
		}
	}
	maxTime := Function.MaxString(timeInterval)
	minTime := Function.MinString(timeInterval)

	update := bson.M{"$set": bson.M{"startMealDay": minTime, "endMealDay": maxTime}}
	mongo.Switches{}.UpdateAll(filter, update)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": "",
	})
}

func EnableOrderSwitch(c *gin.Context) {
	//登录验证
	logier, roleType := UserData{}.isAdminLogin(c)
	if logier == nil || roleType != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "没有权限", "data": "",
		})
		return
	}
	var param map[string]int
	err := c.BindJSON(&param)
	if err != nil || arrays.Contains([]int{0, 1}, param["enable"]) == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": err,
		})
		return
	}
	//查找是否存在
	orderSwitch := mongo.Switches{}.FindOne(bson.M{"name": "order"})
	if orderSwitch == nil {
		s := mongo.Switches{
			ID: primitive.NewObjectID(), Name: "order", CreateTime: time.Now(), Enable: 0, UpdateTime: time.Now(),
		}
		s.CreateRow()
	}
	filter := bson.M{"name": "order"}
	update := bson.M{"$set": bson.M{"enable": param["enable"]}}
	mongo.Switches{}.UpdateAll(filter, update)
	var msg string
	if param["enable"] == 0 {
		msg = "关闭订餐"
	} else {
		msg = "开启订餐"
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg, "data": param["enable"]})
}

func GetOrderSwitch(c *gin.Context) {
	//登录验证
	logier, roleType := UserData{}.isAdminLogin(c)
	if logier == nil || roleType != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "没有权限", "data": "",
		})
		return
	}
	filter := bson.M{"name": "order"}
	res := mongo.Switches{}.FindOne(filter)
	var data bool
	if res["enable"] == int32(1) {
		data = true
	} else {
		data = false
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
}

func (ud UserData) isAdminLogin(c *gin.Context) (interface{}, int32) {
	session := sessions.Default(c)
	logier := session.Get("logier")
	roleType := session.Get("roleType").(int32)
	if logier != "" {
		fmt.Println(roleType)
		return logier, roleType
	} else {
		return "", 0
	}
}
