package yc

const Version = "v0.0.1"

type SignMethod string

const (
	SignMethodUnknown = SignMethod("unknown")
	SignMethodMd5     = SignMethod("md5")
	SignMethodSha1    = SignMethod("sha1")
	SignMethodSha256  = SignMethod("sha256")
	SignMethodSha512  = SignMethod("sha512")
)

func IsValidSignMethod(signMethod SignMethod) bool {
	return signMethod == SignMethodMd5 || signMethod == SignMethodSha1 || signMethod == SignMethodSha256 || signMethod == SignMethodSha512
}
