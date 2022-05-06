package cos

import (
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"os"
)

var Error = struct {
	AccountNotFound   errors.Error `code:"10000" msg:"account not found"`
	AccountExists     errors.Error `code:"10001" msg:"account already exists"`
	IncorrectPassword errors.Error `code:"10002" msg:"incorrect password"`
	AuthorizeFailed   errors.Error `code:"10003" msg:"authorize failed"`
}{}

func init() {
	if err := errors.InitErrorStruct(&Error); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
