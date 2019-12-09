package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xd-meal-back-end/middleware/mongo"
	"github.com/xd-meal-back-end/pkg/gredis"
	"github.com/xd-meal-back-end/pkg/logging"
	"github.com/xd-meal-back-end/pkg/setting"
	"github.com/xd-meal-back-end/pkg/util"
	"github.com/xd-meal-back-end/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	mongo.Setup()
	logging.Setup()
	_ = gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}
