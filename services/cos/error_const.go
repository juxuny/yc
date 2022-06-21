package cos

import (
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"os"
)

var Error = struct {
	AccountNotFound               errors.Error `code:"10000" msg:"account not found"`
	AccountExists                 errors.Error `code:"10001" msg:"account already exists"`
	IncorrectPassword             errors.Error `code:"10002" msg:"incorrect password"`
	AuthorizeFailed               errors.Error `code:"10003" msg:"authorize failed"`
	AccountForbidden              errors.Error `code:"10004" msg:"account forbidden"`
	NoPermissionAccessUserInfo    errors.Error `code:"10005" msg:"no permission access user info"`
	NoDataModified                errors.Error `code:"10006" msg:"no data modified"`
	NamespaceNotFound             errors.Error `code:"10007" msg:"namespace not found"`
	NoPermissionDeleteNamespace   errors.Error `code:"10008" msg:"no permission delete"`
	NamespaceDuplicated           errors.Error `code:"10009" msg:"namespace duplicated"`
	NoPermissionToAssessConfig    errors.Error `code:"10010" msg:"no permission to assess config"`
	ConfigIdDuplicated            errors.Error `code:"10011" msg:"configId duplicated"`
	ConfigNotFound                errors.Error `code:"10012" msg:"config not found"`
	Unauthorized                  errors.Error `code:"10013" msg:"unauthorized"`
	DeleteNotAllowed              errors.Error `code:"10014" msg:"delete not allowed"`
	MissingArguments              errors.Error `code:"10015" msg:"missing arguments"`
	ConfigDisabled                errors.Error `code:"10016" msg:"config disabled"`
	MissingConfigId               errors.Error `code:"10037" msg:"missing configId"`
	KeyDuplicated                 errors.Error `code:"10038" msg:"key duplicated"`
	KeyNotFound                   errors.Error `code:"10039" msg:"key not found"`
	NoPermissionUpdateUserInfo    errors.Error `code:"10040" msg:"not permission to update user info"`
	IllegalUserData               errors.Error `code:"10041" msg:"illegal user data"`
	NotAllowedCreateAccount       errors.Error `code:"10042" msg:"not allowed create account"`
	NoPermissionAccessNamespace   errors.Error `code:"10043" msg:"no permission access namespace"`
	AccessKeyNotFound             errors.Error `code:"10044" msg:"access_key not found"`
	NoPermissionToAccessAccessKey errors.Error `code:"10045" msg:"no permission access"`
	AccessKeyDisabled             errors.Error `code:"10046" msg:"access key disabled"`
}{}

func init() {
	if err := errors.InitErrorStruct(&Error); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
