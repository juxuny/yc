package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/juxuny/yc/env"
	"github.com/juxuny/yc/log"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	verbose = env.IsDebug()
	logger  = log.NewRpcLogger()
)

func GetJSON(url string, out interface{}, headers ...map[string]interface{}) (code int, err error) {
	logger.Debug("get: ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1, fmt.Errorf("request error: %v", err)
	}
	if len(headers) > 0 {
		for _, h := range headers {
			for k, v := range h {
				req.Header.Set(k, fmt.Sprint(v))
			}
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, fmt.Errorf("request error: %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("read body error: %v", err)
	}
	if verbose {
		logger.Debug(string(data))
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return -1, fmt.Errorf("unmrshal json erro: %s", err)
	}
	return resp.StatusCode, nil
}

// PostJSON  POST数据到指定的URL
func PostJSON(url string, param interface{}, out interface{}, headers ...map[string]interface{}) (code int, err error) {
	d, _ := json.Marshal(param)
	logger.Debug("post: ", url, " data: ", string(d))
	body := bytes.NewReader(d)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return -1, fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if len(headers) > 0 {
		for _, h := range headers {
			for k, v := range h {
				req.Header.Set(k, fmt.Sprint(v))
			}
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, fmt.Errorf("request error: %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("read body error: %v", err)
	}
	if verbose {
		logger.Debug(string(data))
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return -1, fmt.Errorf("unmrshal json erro: %s", err)
	}
	return resp.StatusCode, nil
}

func PostJSONWithTimeout(url string, param interface{}, out interface{}, timeout time.Duration, headers ...map[string]interface{}) (code int, err error) {
	d, _ := json.Marshal(param)
	logger.Debug("post: ", url, " data: ", string(d))
	body := bytes.NewReader(d)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return -1, fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if len(headers) > 0 {
		for _, h := range headers {
			for k, v := range h {
				req.Header.Set(k, fmt.Sprint(v))
			}
		}
	}
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("request error: %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("read body error: %v", err)
	}
	if verbose {
		logger.Debug(string(data))
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return -1, fmt.Errorf("unmrshal json erro: %s", err)
	}
	return resp.StatusCode, nil

}
