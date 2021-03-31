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


func main()  {
	/*when you create order succeed, you will recieve a Virtual bank account transaction
	 *User recharge to virtual bank account， after bank success, you will receive notify
	*by your special callback
	 */
	req := transaction.BankTransferInitializeReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.UserRequestIp = "123.123.123.12"
	req.ProductDesc = "test"
	req.CallbackUrl = "http://localhost"
	req.ExpireAt = "10"
	req.UserPhone = "+2348160564736"

	ret, err := transaction.ApiBankTransferInitializeReq(req, mConf)
	if err != nil{
		fmt.Println(ret, err)
	}

	if ret.Code != "00000"{
		return
	}

	{
		reqStatus := transaction.BankTransferStatusReq{Reference: req.Reference}
		ret, err := transaction.ApiBankTransferStatusReq(reqStatus, mConf)
		if err != nil {
			fmt.Println(ret, err)
		}
	}
}