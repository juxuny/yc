package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

var HashHelper = hashHelper{}

type hashHelper struct{}

func (hashHelper) sumFromBytesAsString(h hash.Hash, in ...[]byte) string {
	for _, buf := range in {
		h.Write(buf)
	}
	return fmt.Sprintf("%02x", h.Sum(nil))
}

func (hashHelper) sumFromBytesAsBytes(h hash.Hash, in ...[]byte) []byte {
	for _, buf := range in {
		h.Write(buf)
	}
	return h.Sum(nil)
}

func (hashHelper) sumFromStringAsString(h hash.Hash, in ...string) string {
	for _, buf := range in {
		h.Write([]byte(buf))
	}
	return fmt.Sprintf("%02x", h.Sum(nil))
}

func (hashHelper) sumFromStringAsBytes(h hash.Hash, in ...string) []byte {
	for _, buf := range in {
		h.Write([]byte(buf))
	}
	return h.Sum(nil)
}

func (t hashHelper) Sha1(in ...string) (sum string) {
	return t.sumFromStringAsString(sha1.New(), in...)
}

func (t hashHelper) Sha256(in ...string) string {
	return t.sumFromStringAsString(sha256.New(), in...)
}

func (t hashHelper) Sha512(in ...string) string {
	return t.sumFromStringAsString(sha512.New(), in...)
}

func (t hashHelper) Sha1Bytes(in ...[]byte) []byte {
	return t.sumFromBytesAsBytes(sha1.New(), in...)
}

func (t hashHelper) Sha256Bytes(in ...[]byte) []byte {
	return t.sumFromBytesAsBytes(sha256.New(), in...)
}

func (t hashHelper) Sha512Bytes(in ...[]byte) []byte {
	return t.sumFromBytesAsBytes(sha512.New(), in...)
}

func (t hashHelper) Sha256FromBytesToString(in ...[]byte) string {
	return t.sumFromBytesAsString(sha256.New(), in...)
}

func (t hashHelper) Sha512FromBytesToString(in ...[]byte) string {
	return t.sumFromBytesAsString(sha512.New(), in...)
}

func (t hashHelper) Sha1FromBytesToString(in ...[]byte) string {
	return t.sumFromBytesAsString(sha1.New(), in...)
}

func (t hashHelper) Md5(in ...string) string {
	return t.sumFromStringAsString(md5.New(), in...)
}
