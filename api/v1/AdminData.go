package v1

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/arrays"
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
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	data, err := middleware.ReadMenuExcel(file)
	c.JSON(http.StatusOK, gin.H{"msg": err, "data": data})
}

/**
导入菜单
*/
func ImportMenu(c *gin.Context) {
	var param map[string][]mongo.DishesMongo
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": "参数错误",
		})
		return
	}
	if param["data"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": "空数据",
		})
		return
	}
	for _, v := range param["data"] {
		v.CreateRow()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": "",
	})
}

func EnableOrderSwitch(c *gin.Context) {
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
	filter := bson.M{"name": "order"}
	res := mongo.Switches{}.FindOne(filter)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": res["enable"]})
}
