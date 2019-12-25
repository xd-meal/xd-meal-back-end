package Function

import (
	"fmt"
	"github.com/xd-meal-back-end/pkg/logging"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string) (string, error) {
	timeStart := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		logging.Error("httpGet error1")
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error("httpGet error2")
		return "", err
	}
	timeEnd := time.Now()
	logging.Info(fmt.Sprintf("httpGet : (url: %s code: %s, response body: %s, elaspsed time:%s)", url, resp.Status, string(body), timeEnd.Sub(timeStart)))
	return string(body), nil
}
