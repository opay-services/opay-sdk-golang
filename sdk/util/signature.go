package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func SignatureSHA512(data []byte, secretKey string) (ret string) {
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}

func SignatureSHA3512(data []byte, secretKey string) (ret string)  {
	h := hmac.New(sha3.New512, []byte(secretKey))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}
