package Function

import (
	"github.com/xd-meal-back-end/pkg/logging"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (string, error) {
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

	return string(body), nil
}
