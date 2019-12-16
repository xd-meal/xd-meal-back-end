package v1

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xd-meal-back-end/middleware/mongo"
	"github.com/xd-meal-back-end/pkg/e"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

type UserData struct {
}

func Login(c *gin.Context) {
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
	filter := bson.M{"email": param["email"], "password": fmt.Sprintf("%x", md5.Sum([]byte(param["password"])))}
	info := mongo.FindOneSelected(filter, "meal", "user")
	if info != nil {
		session := sessions.Default(c)
		id, _ := json.Marshal(info["_id"])
		logier, _ := strconv.Unquote(string(id))
		session.Set("logier", logier)
		session.Set("email", param["email"])
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
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}
	var param map[string]string
	err := c.BindJSON(&param)
	if err != nil || param["password"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 417,
			"msg":  e.GetMsg(400),
			"data": "参数不能为空",
		})
		return
	}
	id, _ := primitive.ObjectIDFromHex(logier.(string))
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
			"code": 0,
			"msg":  "请先登录",
			"data": "",
		})
		return
	}

}
