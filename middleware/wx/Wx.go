package wx

import (
	"encoding/json"
	"fmt"
	"github.com/xd-meal-back-end/middleware/mongo"
	"github.com/xd-meal-back-end/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"time"
)

const CorpSecret = "tWk0PLidHSr-jcLJM73EKeQnSUq39RRTHX_AY_-6tIM"
const CorpID = "wxe2be6e5c62e7b072"
const AgentId = 1000033
const RedirectUri = "http://meal.wcmoon.com/api/v1/WeiXinLogin"

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
		client := &http.Client{}
		url2 := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorpID, CorpSecret)
		request, err := http.NewRequest("GET", url2, nil)
		response, err := client.Do(request)
		if err != nil {
			logging.Error(err)
		}
		defer response.Body.Close()
		if response.StatusCode == 200 {
			r, err := ioutil.ReadAll(response.Body)
			if err != nil {
				logging.Error(err)
			}
			acc := map[string]string{}
			_ = json.Unmarshal(r, &acc)
			accessToken = acc["access_token"]
			mongo.AccessToken{}.UpdateAll(bson.M{"name": "accessToken"}, bson.M{"$set": bson.M{"token": accessToken, "updateTime": currentTime}})
		} else {
			logging.Error(fmt.Sprintf("url:%s;return:%v", url2, response))
			accessToken = ""
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
