package wx

import (
	"encoding/json"
	"fmt"
	"github.com/xd-meal-back-end/Function"
	"github.com/xd-meal-back-end/middleware/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const CorpSecret = "tWk0PLidHSr-jcLJM73EKeQnSUq39RRTHX_AY_-6tIM"
const CorpID = "wxe2be6e5c62e7b072"
const AgentId = 1000033
const RedirectUri = "/api/v1/WeiXinLogin"

type WeiXin struct {
}

func (w WeiXin) GetAccessToken() string {
	accessToken := ""
	//一小时前
	currentTime := time.Now()
	h, _ := time.ParseDuration("-1h")
	h1 := currentTime.Add(h)
	exist := mongo.AccessToken{}.FindOne(bson.M{"updateTime": bson.M{"$gte": h1}})
	if exist == nil {

		url2 := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorpID, CorpSecret)
		res, err := Function.HttpGet(url2)
		if err == nil {
			acc := map[string]string{}
			_ = json.Unmarshal([]byte(res), &acc)
			accessToken = acc["access_token"]
			mongo.AccessToken{}.UpdateAll(bson.M{"name": "accessToken"}, bson.M{"$set": bson.M{"token": accessToken, "updateTime": currentTime}})
		}
	} else {
		accessToken = exist["token"].(string)
	}
	return accessToken
}

//func (w *WeiXin)Login() string{
//
//	os.Exit(1)
//	//https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=wxe2be6e5c62e7b072&agentid=1000033&state=wework_redirect_xd&redirect_uri=http%3A%2F%2Fmeal.wcmoon.com%2F
//}
