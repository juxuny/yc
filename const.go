package yc

type SignMethod string

const (
	SignMethodMd5    = SignMethod("md5")
	SignMethodSha1   = SignMethod("sha1")
	SignMethodSha128 = SignMethod("sha128")
	SignMethodSha256 = SignMethod("sha256")
)

func IsValidSignMethod(signMethod SignMethod) bool {
	return signMethod == SignMethodMd5 || signMethod == SignMethodSha1 || signMethod == SignMethodSha128 || signMethod == SignMethodSha256
}
