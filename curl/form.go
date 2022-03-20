package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostForm(u string, values url.Values, out interface{}, headers ...map[string]interface{}) (code int, err error) {
	body := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return -1, errors.Wrap(err, "create post request error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, h := range headers {
		for k, v := range h {
			req.Header.Add(k, fmt.Sprintf("%v", v))
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, errors.Wrap(err, "do request error")
	}
	code = resp.StatusCode
	if resp.StatusCode/100 != 2 {
		// request failed
		err = fmt.Errorf("request failed")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, errors.Wrap(err, "io error")
	}
	err = json.Unmarshal(data, out)
	return
}
