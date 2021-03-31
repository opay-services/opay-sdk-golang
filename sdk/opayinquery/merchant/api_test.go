package merchant

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
	"time"
)

var mConf *conf.OpayMerchantConf
func init()  {
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

func TestApiInfoMerchantReq(t *testing.T) {
	req := InfoMerchantReq{Email:"lijian2233@163.com"}
	resp, err := ApiInfoMerchantReq(req, mConf)
	fmt.Println(resp, err)
}
