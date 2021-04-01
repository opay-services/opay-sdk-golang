package main

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"time"
)

func init() {
	m1 := conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
		"NGN",
	)

	m1.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})
	rand.Seed(time.Now().Unix())

	m2 := conf.InitEnv(
		"OPAYPUB16171838837270.48165639591585996",
		"OPAYPRV16171838837270.964258481459751",
		"qqVzWbylxkjJORvF",
		"256621033118750",
		"sandbox",
		"US",
	)

	m2.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})
	rand.Seed(time.Now().Unix())
}

func main() {
	m1 := conf.GetMerchantConfig("256620112018025")
	m2 := conf.GetMerchantConfig("256621033118750")

	{
		//create cashier order
		req := cashier.CashierInitializeReq{}
		req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
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
		rsp, err := cashier.ApiCashierInitializeReq(req, m1)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rsp.Code != "00000" {
			return
		}

		//query order status
		ret, err := cashier.ApiCashierStatusReq(cashier.CashierStatusReq{Reference: req.Reference}, m1)
		if ret.Code != "00000" {

		}
		//close order
		//only init status can close
		c, err := cashier.ApiCashierCloseReq(cashier.CashierCloseReq{Reference: req.Reference}, m1)
		if c.Code != "00000" {

		}
	}

	{
		{
			//create cashier order
			req := cashier.CashierInitializeReq{}
			req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
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
			rsp, err := cashier.ApiCashierInitializeReq(req, m2)
			if err != nil {
				fmt.Println(err)
				return
			}
			if rsp.Code != "00000" {
				return
			}

			//query order status
			ret, err := cashier.ApiCashierStatusReq(cashier.CashierStatusReq{Reference: req.Reference}, m2)
			if ret.Code != "00000" {

			}
			//close order
			//only init status can close
			c, err := cashier.ApiCashierCloseReq(cashier.CashierCloseReq{Reference: req.Reference}, m2)
			if c.Code != "00000" {

			}
		}
	}
}
