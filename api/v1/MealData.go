package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xd-meal-back-end/middleware"
	"github.com/xd-meal-back-end/middleware/mongo"
	"github.com/xd-meal-back-end/pkg/e"
	"github.com/xd-meal-back-end/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

type MealData struct {
}

func DownloadMenu(c *gin.Context) {
	arr := [][]string{
		{"", "", "", "中餐"},
		{"菜品名称", "菜品明细", "供应商", "菜品编号"},
		{},
		{},
		{},
		{},
		{},
		{},
		{"", "", "", "晚餐"},
		{"菜品名称", "菜品明细", "供应商", "菜品编号"},
	}
	middleware.ExportTmp(arr, c)
}

/**
获取可选菜单表
*/
func GetDishes(c *gin.Context) {
	//登录验证
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}
	//选饭启动区间区间
	filterSwitch := bson.M{"name": "order", "enable": 1}
	switches := mongo.Switches{}.FindOne(filterSwitch)
	if switches == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 401, "msg": "还未开启选餐"})
		return
	}
	filter := bson.M{"mealDay": bson.M{"$gte": switches["startMealDay"], "$lte": switches["endMealDay"]}}
	logging.Info(filter)
	res := mongo.FindAllSelected(filter, "meal", "dishes")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": res,
	})
}

/**
用户点餐
*/
func OrderDishes(c *gin.Context) {
	//登录验证
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}

	var param map[string][]string
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": err,
		})
		return
	}
	dishIds := param["dishIds"]
	idList := make([]primitive.ObjectID, len(dishIds))
	for i, id := range dishIds {
		idList[i], _ = primitive.ObjectIDFromHex(id)
	}
	currentTime := time.Now()
	dishes := mongo.FindAllSelected(bson.M{"_id": bson.M{"$in": idList}}, "meal", "dishes")

	for _, v := range dishes {
		id, _ := json.Marshal(v["_id"])
		dishId, _ := strconv.Unquote(string(id))
		insert := mongo.UserDishes{ID: primitive.NewObjectID(), Uid: logier.(string), DishId: dishId,
			MealDay: v["mealDay"].(string), OrderTime: currentTime, BadEval: 0}
		insert.CreateRow()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": "",
	})
}

func GetOrderDishes(c *gin.Context) {
	//登录验证
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}
	userDishes := mongo.FindAllSelected(bson.M{"uid": logier.(string), "status": 0}, "meal", "userDishes")
	//ArrayColumn
	columns := make([]interface{}, 0, len(userDishes))
	for _, val := range userDishes {
		if v, ok := val["dishId"]; ok {
			objId, _ := primitive.ObjectIDFromHex(v.(string))
			columns = append(columns, objId)
		}
	}
	dishes := mongo.FindAllSelected(bson.M{"_id": bson.M{"$in": columns}}, "meal", "dishes")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": dishes,
	})
}

func UpdateUserOrder(c *gin.Context) {
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}
	dish := mongo.UserDishes{}
	//objId, _ := primitive.ObjectIDFromHex("5dea01d126a606122cf74d8b")
	filter := bson.M{"uid": "de4db7a7cfb83e4f6a61a25"}
	update := bson.M{
		"$set": bson.M{"status": 3},
	}
	success := dish.UpdateAll(filter, update)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": success,
	})
}
