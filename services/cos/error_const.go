package cos

import (
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"os"
)

var Error = struct {
	AccountNotFound            errors.Error `code:"10000" msg:"account not found"`
	AccountExists              errors.Error `code:"10001" msg:"account already exists"`
	IncorrectPassword          errors.Error `code:"10002" msg:"incorrect password"`
	AuthorizeFailed            errors.Error `code:"10003" msg:"authorize failed"`
	AccountForbidden           errors.Error `code:"10004" msg:"account forbidden"`
	NoPermissionAccessUserInfo errors.Error `code:"10005" msg:"no permission access user info"`
	NoDataModified             errors.Error `code:"10006" msg:"no data modified"`
	NamespaceNotFound          errors.Error `code:"10007" msg:"namespace not found"`
}{}

func init() {
	if err := errors.InitErrorStruct(&Error); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
