package curl

import "fmt"

type ContentType string

const (
	ContentTypeJson = "application/json"
)

func Post(url string, contentType ContentType, data interface{}, out interface{}, header ...map[string]interface{}) (int, error) {
	if contentType == ContentTypeJson {
		return PostJSON(url, data, out, header...)
	}
	return 0, fmt.Errorf("unknown content-type: %v", contentType)
}
