package transaction

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
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

func TestApiSupportBanksReq(t *testing.T) {
	ret, err := ApiSupportBanksReq(SupportBanksReq{Country:"NG"})
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiByBankCardReq(t *testing.T) {
	req := ByBankCardReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.BankCode = "103"
	req.Reason = "test"
	req.CustomerEmail = "xxx@163.com"
	req.Return3dsUrl = "http://localhost:8080"
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.FirstName = "li"
	req.LastName = "jian"
	req.CardCVC = "561"
	req.CardDateYear = "22"
	req.CardDateMonth = "12"
	req.CardNumber = "5061460410121111105"

	ret, err := ApiByBankCardReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiByBankAccountReq(t *testing.T) {
	req := ByBankAccountReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.BankAccountNumber = "2070043686"
	req.BankCode = "033"
	req.Reason = "test"
	req.CustomerPhone = "+2348160564736"
	req.Return3dsUrl = "http://localhost:8080"
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.Bvn = "1234567890"
	req.DobYear = "2000"
	req.DobMonth = "10"
	req.DobDay = "30"

	ret, err := ApiByBankAccountReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiStatusReq(t *testing.T) {
	req := StatusReq{Reference:"testlijian_1616747803824832000"}
	ret, err := ApiStatusReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiUssdCodeReq(t *testing.T) {
	req := UssdCodeReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.BankCode = "033"
	req.UserRequestIp = "123.123.123.122"
	req.ExpireAt = "10"
	req.ProductDesc = "test"
	req.CallbackUrl = "localhost://localhost:8080"
	req.UserPhone = "+2348160564736"

	ret, err := ApiUssdCodeReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiUssdOrderStatusReq(t *testing.T) {
	req := UssdOrderStatusReq{Reference:"testlijian_1616753993664468000"}
	ret, err := ApiUssdOrderStatusReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiBankTransferInitializeReq(t *testing.T) {
	req := BankTransferInitializeReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.UserRequestIp = "123.123.123.12"
	req.ProductDesc = "test"
	req.CallbackUrl = "http://localhost"
	req.ExpireAt = "10"
	req.UserPhone = "+2348160564736"

	ret, err := ApiBankTransferInitializeReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiBankTransferStatusReq(t *testing.T) {
	req := BankTransferStatusReq{Reference:"testlijian_1616751184699154000"}
	ret, err := ApiBankTransferStatusReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiInputDobReq(t *testing.T) {

}

func TestApiInputOtpReq(t *testing.T) {
	req := InputOtpReq{}
	req.Reference = "testlijian_1616743592134162000"
	req.Otp = "543210"
	ret, err := ApiInputOtpReq(req)
	if err != nil{
		fmt.Println(ret, err)
	}
}

func TestApiInputPinReq(t *testing.T) {
	req := InputPinReq{Pin:"1105"}
	req.Reference = "testlijian_1616746186822213000"
	ret, err := ApiInputPinReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}

func TestApiInputPhoneReq(t *testing.T) {
	req := InputPhoneReq{Phone:"+2348160564736"}
	req.Reference = "testlijian_1616746356023949000"
	ret, err := ApiInputPhoneReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}

