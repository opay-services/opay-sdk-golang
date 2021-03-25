package main

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
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


	/*conf.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})*/
	rand.Seed(time.Now().Unix())
}


func main()  {
	//create cashier order
	req := cashier.CashierInitializeReq{}
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.MchShortName = "lijian_test"
	req.ProductName = "goods"
	req.ProductDesc = "test goods"
	req.UserPhone = "+2349876543210"
	req.UserRequestIp = "123.123.123.123"
	req.Amount = "100"
	req.Currency = "NGN"
	req.PayTypes = []string{"BalancePayment", "BonusPayment", "OWealth"}
	req.PayMethods = []string{"account", "qrcode", "bankCard", "bankAccount"}
	req.ExpireAt = 10
	req.CallbackUrl = "http://localhost:8080"
	req.ReturnUrl = "http://localhost:8080"
	rsp, err := cashier.ApiCashierInitializeReq(req)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("ret:%+v, err:%v\n", rsp, err)

	//query order status
	ret, err := cashier.ApiCashierStatusReq(cashier.CashierStatusReq{Reference:req.Reference})
	fmt.Printf("ret:%+v, err:%v\n", ret, err)

	//close order
	//only init status can close
	c, err := cashier.ApiCashierCloseReq(cashier.CashierCloseReq{Reference:req.Reference})
	fmt.Printf("ret:%+v, err:%v\n", c, err)
}
