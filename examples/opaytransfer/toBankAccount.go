package main

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/transfer"
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
	//create a tarnsfer
	req := transfer.ToBankReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = transfer.BankReceiver{
		BankCode:"033",
		BankAccountNumber:"2070043686",
		NameCheck:"yes",
		Name:"test",
	}

	rsp, err := transfer.ApiTransferToBank(req, mConf)
	if err != nil{
		fmt.Println(err)
		return
	}
	if rsp.Code != "00000"{
		return
	}

	//query status
	//The transfer will be successful after a few minutes, depending on the processing time of the bank
	for i:=0; i<10; i++ {
		rsp, err = transfer.ApiStatusToBankReq(transfer.StatusToBankReq{Reference: req.Reference}, mConf)
		if err != nil{
			continue
		}

		if rsp.Code == "00000" && rsp.Data.Status == "SUCCESSFUL"{
			break
		}
	}
}