package v1

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
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
	var data bool
	if res["enable"] == int32(1) {
		data = true
	} else {
		data = false
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
}

func GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0, "msg": "邮箱不能为空", "data": "",
		})
		return
	}
	filter := bson.M{"email": email}
	res := mongo.UserMongo{}.FindOne(filter)
	data := map[string]interface{}{"id": res["_id"], "name": res["name"].(string), "email": res["email"].(string)}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
}

func AddMenuSingle(c *gin.Context) {
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil || param["uid"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0, "msg": "用户id不能为空", "data": "",
		})
		return
	}
	uid := param["uid"]

	res := mongo.DishesMongo{}.GetOptionalBuffet()
	currentTime := time.Now()
	for _, v := range res {
		id, _ := json.Marshal(v["_id"])
		dishId, _ := strconv.Unquote(string(id))
		insert := mongo.UserDishesMongo{ID: primitive.NewObjectID(), Uid: uid, DishId: dishId, Name: v["name"].(string), Dsc: v["dsc"].(string), Supplier: v["supplier"].(string), MealNum: v["mealNum"].(int32),
			TypeA: v["typeA"].(int32), MealDay: v["mealDay"].(string), OrderTime: currentTime, BadEval: false}

		insert.CreateRow()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": "",
	})
}

func GetMealTotal(c *gin.Context) {
	res := mongo.UserDishesMongo{}.GetTotalByOrdered()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": res,
	})
}

/**
后台重置密码
*/
func ResetPassword(c *gin.Context) {
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil || param["uid"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0, "msg": "用户id不能为空", "data": "",
		})
		return
	}
	uid := param["uid"]
	mongo.UserMongo{}.ResetPassWord(uid)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "重置密码成功",
		"data": nil,
	})
}
