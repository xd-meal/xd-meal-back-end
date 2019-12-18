package v1

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xd-meal-back-end/middleware/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

type UserData struct {
}

type RequestEvalDish struct {
	ID   string `json:"id" form:"name"`
	Eval bool   `json:"eval" form:"eval"`
}

func Login(c *gin.Context) {
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  err,
			"data": "",
		})
		return
	}
	filter := bson.M{"email": param["email"], "password": fmt.Sprintf("%x", md5.Sum([]byte(param["password"])))}
	info := mongo.FindOneSelected(filter, "meal", "user")
	if info != nil {
		session := sessions.Default(c)
		id, _ := json.Marshal(info["_id"])
		logier, _ := strconv.Unquote(string(id))
		session.Set("logier", logier)
		session.Set("email", param["email"])
		session.Set("roleType", info["type"])
		_ = session.Save()
		fmt.Println("logier:", session.Get("logier"))
		fmt.Println("email:", session.Get("email"))
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录失败"})
	}
}

func LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("logier")
	session.Delete("email")
	_ = session.Save()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "退出成功", "data": ""})
}

func CheckUserLogin(c *gin.Context) {
	logier := UserData{}.isLogin(c)
	if logier != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "已登陆", "data": logier})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "未登录"})
	}
}

func (ud UserData) isLogin(c *gin.Context) interface{} {
	session := sessions.Default(c)
	logier := session.Get("logier")
	if logier != nil {
		return logier
	} else {
		return nil
	}
}

func ResetPasswordByUser(c *gin.Context) {
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
		})
		return
	}
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil || param["password"] == "" || param["oldPassword"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  "参数不能为空",
			"data": "",
		})
		return
	}
	id, _ := primitive.ObjectIDFromHex(logier.(string))
	checkUser := mongo.UserMongo{}.FindOne(bson.M{"_id": id, "password": fmt.Sprintf("%x", md5.Sum([]byte(param["oldPassword"])))})
	if checkUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  "原密码错误",
			"data": "",
		})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"password": fmt.Sprintf("%x", md5.Sum([]byte(param["password"])))}}
	mongo.UserMongo{}.UpdateAll(filter, update)
	session := sessions.Default(c)
	session.Delete("logier")
	session.Delete("email")
	_ = session.Save()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "密码修改成功，请重新登录"})
}

func EvalDish(c *gin.Context) {
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
		})
		return
	}
	param := RequestEvalDish{}
	if c.ShouldBind(&param) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417, "msg": "参数错误", "data": "",
		})
		return
	}
	id, _ := primitive.ObjectIDFromHex(param.ID)
	uid := logier
	filter := bson.M{"_id": id, "uid": uid}
	dish := mongo.UserDishesMongo{}.FindOne(filter)
	if dish == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "要评价的订单不存在"})
		return
	}
	update := bson.M{"$set": bson.M{"badEval": param.Eval}}
	mongo.UserDishesMongo{}.UpdateAll(filter, update)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "评价成功"})
}

func GetDishCode(c *gin.Context) {
	logier := UserData{}.isLogin(c)
	if logier == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "请先登录", "data": "",
		})
		return
	}
	currentTime := time.Now()
	breakfastStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 8, 30, 0, 0, currentTime.Location())
	breakfastEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 10, 0, 0, 0, currentTime.Location())

	lunchStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 11, 50, 0, 0, currentTime.Location())
	lunchEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 13, 30, 0, 0, currentTime.Location())

	dinnerStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 17, 50, 0, 0, currentTime.Location())
	dinnerEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 19, 30, 0, 0, currentTime.Location())
	token := "100000000000000000000000"
	t := ""
	if currentTime.Before(breakfastEndTime) && currentTime.After(breakfastStartTime) {
		t = "A"
		token = logier.(string)
	} else if currentTime.Before(lunchEndTime) && currentTime.After(lunchStartTime) {
		t = "B"
		filter := bson.M{"mealDay": time.Now().Format("2006-01-02"), "typeA": 1}
		res := mongo.UserDishesMongo{}.FindOne(filter)
		id, _ := json.Marshal(res["_id"])
		token, _ = strconv.Unquote(string(id))
	} else if currentTime.Before(dinnerEndTime) && currentTime.After(dinnerStartTime) {
		t = "C"
		filter := bson.M{"mealDay": time.Now().Format("2006-01-02"), "typeA": 2}
		res := mongo.UserDishesMongo{}.FindOne(filter)
		id, _ := json.Marshal(res["_id"])
		token, _ = strconv.Unquote(string(id))
	} else {
		t = "D"
		token = "100000000000000000000002"
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": len(token), "data": t + token})
}

func ScanDishCode(c *gin.Context) {
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil || param["token"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  "参数不能为空",
			"data": "",
		})
		return
	}
	t := param["token"][0:1]
	fmt.Println(t)
	token := param["token"][1:len(param["token"])]
	fmt.Println(token)
	update := bson.M{"$set": bson.M{"status": 1}}
	switch t {
	case "A":
		//早餐
		filter := bson.M{"uid": token, "mealDay": time.Now().Format("2006-01-02"), "typeA": 3}
		userDish := mongo.UserDishesMongo{}.FindOne(filter)
		if userDish == nil {
			insert := mongo.UserDishesMongo{ID: primitive.NewObjectID(), Uid: token, DishId: "", Name: "早餐", Supplier: "心动食堂", Status: 1,
				TypeA: 3, MealDay: time.Now().Format("2006-01-02"), OrderTime: time.Now(), BadEval: false}
			insert.CreateRow()
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "早餐扫码成功", "data": ""})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "不能重复取餐", "data": ""})
		}
		return
	case "B":
		//午餐
		id, _ := primitive.ObjectIDFromHex(token)
		filter := bson.M{"_id": id}
		up := mongo.UserDishesMongo{}.UpdateAll(filter, update)
		userDish := mongo.UserDishesMongo{}.FindOne(filter)
		if up == int64(1) {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "取餐成功", "data": userDish})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "不能重复取餐", "data": userDish})
		}
		return
	case "C":
		id, _ := primitive.ObjectIDFromHex(token)
		filter := bson.M{"_id": id}
		up := mongo.UserDishesMongo{}.UpdateAll(filter, update)
		userDish := mongo.UserDishesMongo{}.FindOne(filter)
		if up == int64(1) {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "取餐成功", "data": userDish})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "不能重复取餐", "data": userDish})
		}
		return
		//晚餐
	default:
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "未到取餐时间", "data": ""})
		return
	}
}
