package conf

import (
	"sync"
)

//https://open.opayweb.com/settings-->API keys
type opayMerchantConf struct {
	PublicKey         string
	SecretKey         string
	AesKey            string
	Id                string //Merchant ID
	Env               string
	Log               func(a ...interface{})
}

var once = &sync.Once{}
var conf *opayMerchantConf

func InitEnv(publicKey, secretKey, aesKey, id, env string) {
	once.Do(func() {
		conf = &opayMerchantConf{Id: id, PublicKey: publicKey, SecretKey: secretKey, AesKey: aesKey}
		if env != "live" {
			conf.Env = "http://sandbox-cashierapi.opayweb.com"
		} else {
			conf.Env = "https://cashierapi.opayweb.com"
		}
	})
}

func GetApiHost() string {
	return conf.Env
}

func GetMerchantId() string {
	return conf.Id
}

func GetPublicKey() string {
	return conf.PublicKey
}

func GetSecretKey() string {
	return conf.SecretKey
}

func GetAesKey() string {
	return conf.AesKey
}

func GetLog() (func(a...interface{}))  {
	return conf.Log
}

func SetLog(f func(a ...interface{}))  {
	conf.Log = f
}

