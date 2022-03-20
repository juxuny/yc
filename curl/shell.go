package curl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func wrap(s string, wrapper ...string) string {
	w := "\""
	if len(wrapper) > 0 {
		w = wrapper[0]
	}
	return w + s + w
}

func GenShellCommand(method string, path string, header http.Header, bodyValues url.Values) string {
	allowHeader := map[string]bool{
		"Cookie":       true,
		"Accept":       true,
		"Content-Type": true,
		"User-Agent":   true,
		"Host":         true,
	}
	shell := []string{
		"curl",
	}
	if method == http.MethodPost {
		shell = append(shell, "-d")
		shell = append(shell, wrap(bodyValues.Encode(), "'"))
	}
	for k, v := range header {
		if !allowHeader[k] {
			continue
		}
		for i := 0; i < len(v); i++ {
			shell = append(shell, "-H", wrap(fmt.Sprintf("%s: %v", k, v[i]), "'"))
		}
	}
	shell = append(shell, wrap(path, "'"))
	return strings.Join(shell, " ")
}

func GenShellCommandWithBody(method string, path string, header http.Header, body []byte) string {
	values := url.Values{}
	if method == http.MethodPost {
		m := map[string]interface{}{}
		_ = json.Unmarshal(body, &m)
		for k, v := range m {
			values.Add(k, fmt.Sprintf("%v", v))
		}
	}
	return GenShellCommand(method, path, header, values)
}
