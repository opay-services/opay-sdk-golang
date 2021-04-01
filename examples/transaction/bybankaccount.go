package main

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/transaction"
	"math/rand"
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


func main(){
	//input-opt, input-pin, input-phone, input-dob reference bybankcard.go
	req := transaction.ByBankAccountReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.BankAccountNumber = "2070043686"
	req.BankCode = "033"
	req.Reason = "test"
	req.CustomerPhone = "+2348160564736"
	req.Return3dsUrl = "http://localhost:8080"
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.Bvn = "1234567890"
	req.DobYear = "2000"
	req.DobMonth = "10"
	req.DobDay = "30"

	ret, err := transaction.ApiByBankAccountReq(req, mConf)
	if err != nil{
		fmt.Println(ret, err)
	}

	//query status
	{
		reqStatus := transaction.StatusReq{Reference:req.Reference}
		ret, err := transaction.ApiStatusReq(reqStatus, mConf)
		if err != nil{
			fmt.Println(ret, err)
		}
	}
}