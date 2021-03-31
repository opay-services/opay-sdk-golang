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
	//create a transfer
	req := transfer.ToOWalletMerchantReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = transfer.OWalletReceiverMerchant{
		MerchantId:"256620112018024",
		Type:"MERCHANT",
		Name:"Andy Lee",
	}
	rsp, err := transfer.ApiTransferToOWalletMerchant(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	if rsp.Code != "00000"{
		return
	}

	//query status
	for i:=0; i<5; i++ {
		rsp, err = transfer.ApiStatusToWalletReq(transfer.StatusToWalletReq{Reference: req.Reference})
		if err != nil{
			fmt.Println(err)
			continue
		}

		if rsp.Code == "00000" && rsp.Data.Status == "SUCCESSFUL"{
			break
		}
		time.Sleep(2*time.Second)
	}
}