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
	/*create ussd order, if succeed, you will get ussd code.
	 *then, dail ussd code, Follow the instructions to complete the transaction
	 *if transaction succeed, we will notify callback.
	 *
	 *Of course, Of course, you can also check the status of the order regularly until the final
	 *status of the order。SUCCESS, FAIL is final status
	 */
	req := transaction.UssdCodeReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.BankCode = "033"
	req.UserRequestIp = "123.123.123.122"
	req.ExpireAt = "10"
	req.ProductDesc = "test"
	req.CallbackUrl = "localhost://localhost:8080"
	req.UserPhone = "+2348160564736"

	ret, err := transaction.ApiUssdCodeReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}

	if ret.Code != "00000"{
		return
	}

	{
		reqStatus := transaction.UssdOrderStatusReq{Reference:req.Reference}
		ret, err := transaction.ApiUssdOrderStatusReq(reqStatus)
		if err != nil{
			fmt.Println(ret, err)
		}

	}
}


