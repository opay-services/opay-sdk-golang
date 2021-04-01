package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/ips"
	"math/rand"
	"os"
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


func callback()  {
	r := gin.Default()

	r.POST("/callback", func(c *gin.Context) {
		buf := make([]byte, 1024)

		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))

		notify := ips.MerchantAcquiring{}
		err := json.Unmarshal(buf[:n], &notify)
		if err != nil {
			fmt.Println(err)
		}else {
			notify.VerfiySignature(mConf)
		}
	})
	r.Run(":8080")
}

func main()  {
	go callback()
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
	req.CallbackUrl = "https://6f237770df1b.ngrok.io/callback"
	req.ReturnUrl = "http://localhost:8080"
	rsp, err := cashier.ApiCashierInitializeReq(req, mConf)
	if err != nil{
		fmt.Println(err)
		return
	}

	if rsp.Code != "00000"{
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
	ret, err := cashier.ApiCashierStatusReq(statusReq, mConf)


	if (ret.Code != "00000"){
		return
	}

	if ret.Data.Status != "SUCCESS"{
		return
	}

	//only success status order can refund
	{
		//refund to bank account
		req := cashier.CashierRefundByBankAccountReq{}
		req.OriginalReference = "testlijian_1616595268369294000"
		req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
		req.Amount = "1"
		req.Currency = "NGN"
		req.RefundType = "refund2bankaccount"
		req.Reason = "test"
		req.BankCode = "033"
		req.Country = "NG"
		req.BankAccountNumber = "2070043686"

		rsp, err :=cashier.ApiCashierRefundByBankAccountReq(req, mConf)
		fmt.Println(fmt.Sprintf("rsp:%+v, err:+%v", rsp, err))
	}

	{
		//refund to opay merchant account
		req := cashier.CashierRefundByOpayMerchantAccountReq{}
		req.OriginalReference = "testlijian_1616595268369294000"
		req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
		req.Amount = "1"
		req.Currency = "NGN"
		req.RefundType = "refund2opayaccount"
		req.Reason = "test"
		req.Country = "NG"
		req.Receiver = cashier.MerchantReceiver{
			MerchantId:"256620112018024",
			Type:"MERCHANT",
		}
		rsp, err :=cashier.ApiCashierRefundByOpayMerchantAccountReq(req, mConf)
		fmt.Println(fmt.Sprintf("rsp:%+v, err:+%v", rsp, err))
	}

	{
		//refund to opay user account
		req := cashier.CashierRefundByOpayUserAccountReq{}
		req.OriginalReference = "testlijian_1616595268369294000"
		req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
		req.Amount = "1"
		req.Currency = "NGN"
		req.RefundType = "refund2opayaccount"
		req.Reason = "test"
		req.Country = "NG"
		req.Receiver = cashier.UserReceiver{
			PhoneNumber:"+2348160564736",
			Type:"USER",
		}

		rsp, err :=cashier.ApiCashierRefundByOpayUserAccountReq(req, mConf)
		fmt.Println(fmt.Sprintf("rsp:%+v, err:+%v", rsp, err))
	}

	{
		//refund to origin pay method
		req := cashier.CashierRefundByOriginReq{}
		req.OriginalReference = "testlijian_1616595268369294000"
		req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
		req.Amount = "1"
		req.Currency = "NGN"
		req.RefundType = "refundoriginal"
		req.Reason = "test"
		req.Country = "NG"

		rsp, err := cashier.ApiCashierRefundByOriginReq(req, mConf)
		fmt.Println(fmt.Sprintf("rsp:%+v, err:+%v", rsp, err))
	}
}
