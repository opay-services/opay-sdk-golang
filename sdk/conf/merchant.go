package conf

import (
	"fmt"
	"strings"
	"sync"
)

//https://open.opayweb.com/settings-->API keys
type OpayMerchantConf struct {
	publicKey  string
	secretKey  string
	aesKey     string
	merchantId string //Merchant ID
	host       string
	currency   string
	Log        func(a ...interface{})
}

var merchantIdConfMap = map[string]*OpayMerchantConf{}
var currencyConfMap = map[string]*OpayMerchantConf{}

var mutex = sync.Mutex{}

func InitEnv(publicKey, secretKey, aesKey, merchantId, env, currency string) *OpayMerchantConf {
	if currency == "" {
		currency = "NGN"
	}

	publicKey = strings.Trim(publicKey, " ")
	secretKey = strings.Trim(secretKey, " ")
	aesKey = strings.Trim(aesKey, " ")
	merchantId = strings.Trim(merchantId, " ")

	if len(publicKey) == 0 || len(secretKey) == 0 || len(aesKey) == 0 || len(merchantId) == 0 {
		panic("init env param publicKey, secretKey, aesKey, id has empty string")
	}

	currency = strings.ToUpper(currency)
	mutex.Lock()
	defer mutex.Unlock()
	m, ok := merchantIdConfMap[merchantId]
	if !ok {
		c := OpayMerchantConf{publicKey: publicKey, secretKey: secretKey, merchantId: merchantId, aesKey: aesKey, currency: currency}
		if env != "live" {
			c.host = "http://sandbox-cashierapi.opayweb.com"
		} else {
			c.host = "https://cashierapi.opayweb.com"
		}
		merchantIdConfMap[merchantId] = &c
		m = &c
	} else {
		fmt.Sprintf("merchant id:%s exsit", merchantId)
	}

	currencyConfMap[currency] = m
	return m
}

func GetMerchantConfigByMerchantId(merchantId string) *OpayMerchantConf {
	m, _ := merchantIdConfMap[merchantId]
	return m
}

func GetMerchantConfigByCurrency(currency string) *OpayMerchantConf  {
	m, _ := currencyConfMap[strings.ToUpper(currency)]
	return m
}

func (m *OpayMerchantConf) GetApiHost() string {
	return m.host
}

func (m *OpayMerchantConf) GetMerchantId() string {
	return m.merchantId
}

func (m *OpayMerchantConf) GetPublicKey() string {
	return m.publicKey
}

func (m *OpayMerchantConf) GetSecretKey() string {
	return m.secretKey
}

func (m *OpayMerchantConf) GetAesKey() string {
	return m.aesKey
}

func (m *OpayMerchantConf) GetLog() (func(a ...interface{})) {
	return m.Log
}

func (m *OpayMerchantConf) SetLog(f func(a ...interface{})) {
	m.Log = f
}
