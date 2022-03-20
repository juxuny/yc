package curl

import (
	"net/http"
	"net/url"
	"testing"
)

func TestGenShellCommand(t *testing.T) {
	s := GenShellCommand(http.MethodGet, "https://dev.backend.hengyangtai.xyz/admin/user/getlist.action?pageNumber=1&pageSize=10&searchText=", http.Header{
		"Cookie": []string{"has_login=true; adminId=129; admin_name=prod; admin_session=b8602c51183845a9b3be5a41ef94d694"},
	}, nil)
	t.Log(s)

	param := url.Values{}
	param.Add("version", "1.0.0")
	s = GenShellCommand(http.MethodPost, "https://dev.backend.hengyangtai.xyz/admin/user/getlist.action?pageNumber=1&pageSize=10&searchText=", http.Header{
		"Cookie": []string{"has_login=true; adminId=129; admin_name=prod; admin_session=b8602c51183845a9b3be5a41ef94d694"},
	}, param)
	t.Log(s)
}
