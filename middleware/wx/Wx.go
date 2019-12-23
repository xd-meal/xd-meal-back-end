package wx

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/xd-meal-back-end/pkg/logging"
	"net/http"
	"time"
)

const CorpSecret = "tWk0PLidHSr-jcLJM73EKeQnSUq39RRTHX_AY_-6tIM"
const CorpID = "wxe2be6e5c62e7b072"
const AgentId = 1000033

func GetAccessToken() string {
	goCache := cache.New(10*time.Minute, 10*time.Minute)
	accessToken, found := goCache.Get("accessToken")
	if found == false {
		response, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorpID, CorpSecret))
		if err != nil {
			logging.Error("accessToken err:" + err.Error())
		}
		accessToken = response
	}
	return accessToken.(string)
}
