package main

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/transaction"
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
	ret, err := transaction.ApiSupportBanksReq(transaction.SupportBanksReq{Country:"NG"})
	if err != nil{
		fmt.Println(ret, err)
	}
}
