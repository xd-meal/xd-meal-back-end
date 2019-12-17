package v1

import (
	"crypto/md5"
	"encoding/json"
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
	"strconv"
	"time"
)

/**
导入外部用户
*/
func ImportUser(c *gin.Context) {
	//登录验证
	logier := UserData{}.isAdminLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
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
			PassWord: fmt.Sprintf("%x", md5.Sum([]byte(v.PassWord))), Type: v.Type, Depart: v.Depart, CreateTime: v.CreateTime}
		insert.CreateRow()
	}
	c.JSON(http.StatusOK, gin.H{"msg": err, "data": data})
}

/**
预览菜单
*/
func ReadMenu(c *gin.Context) {
	//登录验证
	logier := UserData{}.isAdminLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
		})
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
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
	logier := UserData{}.isAdminLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
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
	logier := UserData{}.isAdminLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
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
	logier := UserData{}.isAdminLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
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

func AdminLogin(c *gin.Context) {
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": err,
		})
		return
	}
	filter := bson.M{"email": param["email"], "type": 1, "password": fmt.Sprintf("%x", md5.Sum([]byte(param["password"])))}
	info := mongo.FindOneSelected(filter, "meal", "user")
	if info != nil {
		session := sessions.Default(c)
		id, _ := json.Marshal(info["_id"])
		logier, _ := strconv.Unquote(string(id))
		session.Set("admin_logier", logier)
		session.Set("admin_email", param["email"])
		_ = session.Save()
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录失败"})
	}
}

func AdminLoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("admin_logier")
	session.Delete("admin_email")
	_ = session.Save()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "退出成功", "data": ""})
}

func (ud UserData) isAdminLogin(c *gin.Context) interface{} {
	session := sessions.Default(c)
	logier := session.Get("admin_logier")
	if logier != nil {
		return logier
	} else {
		return nil
	}
}
