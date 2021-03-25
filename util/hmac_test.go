package util

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	"testing"
)

func TestHMacSha512(t *testing.T) {
	fmt.Println(HMacSha512("{\"name\":\"lijian\"}", "OPAYPRV16058646510230.34019403186305675"))
	/*conf.InitEnv(
		"OPAYPUB16058657924070.5531792522369983",
		"OPAYPRV16058657924070.7523622718657195",
		"MzrcwA0sr3TUAYDn",
		"256620112018224",
		"live")*/

	conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
		)


	req := cashier.CashierInitializeReq{}
	req.Reference = "testlijian_12312131213"
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

	cashier.ApiCashierInitializeReq(req)

}
