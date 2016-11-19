package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	sDec, _ := base64.StdEncoding.DecodeString(s)
	return string(sDec)
}
