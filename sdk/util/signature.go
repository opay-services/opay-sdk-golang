package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"golang.org/x/crypto/sha3"
)

func SignatureSHA512(data []byte) (ret string) {
	h := hmac.New(sha512.New, []byte(conf.GetSecretKey()))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

func SignatureSHA3512(data []byte) (ret string)  {
	h := hmac.New(sha3.New512, []byte(conf.GetSecretKey()))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}