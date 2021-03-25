package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HMacSha512(data string, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

