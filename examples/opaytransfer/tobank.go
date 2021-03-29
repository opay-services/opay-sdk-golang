package main

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/transfer"
	"math/rand"
	"time"
)

func init()  {
	conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
	)


	conf.SetLog(func(a ...interface{}) {
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

	rsp, err := transfer.ApiTransferToBank(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	if rsp.Code != "00000"{
		return
	}

	//query status
	rsp, err = transfer.ApiStatusToBankReq(transfer.StatusToBankReq{Reference:req.Reference})
}