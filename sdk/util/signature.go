package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
)

func Signature(data []byte) (ret string) {
	h := hmac.New(sha512.New, []byte(conf.GetSecretKey()))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}