package charge

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
	"time"
)

var mConf *conf.OpayMerchantConf
func init() {
	mConf = conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
		"NGN",
	)

	mConf.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})
	rand.Seed(time.Now().Unix())
}

func TestApiInitializeReq(t *testing.T) {
	req := InitializeReq{}
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.UserRequestIp = "123.123.123.123"
	req.Amount = "100"
	req.Currency = "NGN"
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.ReturnUrl = "http://localhost:8080"
	req.ChargerType = "USER"
	req.ChargerId = "156619102400201625"

	resp, err := ApiInitializeReq(req, mConf)
	fmt.Println(resp, err)
}

func TestApiStatusReq(t *testing.T) {
	req := StatusReq{}
	req.OrderNo = new(string)
	*req.OrderNo = "210331280540251542"
	resp, err := ApiStatusReq(req, mConf)
	fmt.Println(resp, err)
}