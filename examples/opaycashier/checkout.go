package main

import (
	"bufio"
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"os"
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
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.ReturnUrl = "http://localhost:8080"
	rsp, err := cashier.ApiCashierInitializeReq(req)
	if err != nil{
		fmt.Println(err)
		return
	}

	if rsp.Code != "00000" {
		return
	}
	//click rsp.Data.CashierUrl, you can select opay owealth, by guild, then input pincode is:123456
	// otp is:123456。if success,  after 5s will jump req.ReturnUrl web page


	fmt.Println("please press enter...")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Println(input, err)
	}

	//query order status
	statusReq := cashier.CashierStatusReq{Reference:req.Reference}
	ret, err := cashier.ApiCashierStatusReq(statusReq)
	if err != nil{
		fmt.Println(err)
		return
	}

	if ret.Code != "00000"{
		return
	}

}