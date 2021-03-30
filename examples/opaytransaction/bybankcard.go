package main

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/transaction"
	"math/rand"
	"time"
)

func init() {
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

func main() {
	/*bank card transaction with pin + opt
	1. create an order
	2. query order status, if order status is pending,
	   You should keep querying the order status
	   until the final status or the required
	   input status (input-pin, input-opt, input-dob, input-phone)
	3.if return input-status,
	*/
	req := transaction.ByBankCardReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.BankCode = "033"
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

	ret, err := transaction.ApiByBankCardReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}

labstatus:
	//query status
	{
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			reqStatus := transaction.StatusReq{Reference: req.Reference}
			ret, err := transaction.ApiStatusReq(reqStatus)
			if err != nil {
				//you can retry if occur  Network failure
				fmt.Println(ret, err)
				return
			}

			if ret.Code != "00000" {
				return
			}

			switch ret.Data.Status {
			case "INITIAL", "PENDING":
				break
			case "INPUT-PIN":
				goto labpin
			case "INPUT-OTP":
				goto labopt
			case "INPUT-PHONE", "INPUT-DOB", "3DSECURE":
				//realized
				return

			case "SUCCESS", "FAIL", "CLOSE":
				return
			}
		}
		//time out
		return
	}

labpin:
	//input pin
	{

		reqPin := transaction.InputPinReq{Pin: "1105"}
		reqPin.Reference = req.Reference
		ret, err := transaction.ApiInputPinReq(reqPin)
		if err != nil {
			fmt.Println(ret, err)
		}

		if ret.Code != "00000" {
			return
		}

		if ret.Data.Status == "PENDING" {
			goto labstatus
		}
		goto labover
	}

labopt:
	{
		reqOpt := transaction.InputOtpReq{}
		reqOpt.Reference = req.Reference
		reqOpt.Otp = "543210"
		ret, err := transaction.ApiInputOtpReq(reqOpt)
		if err != nil {
			fmt.Println(ret, err)
		}

		if ret.Code != "00000" {
			return
		}

		if ret.Data.Status == "PENDING" {
			goto labstatus
		}
		goto labover
	}

labover:
	//query status, if order status is "SUCCESS" you will get token, you can transaction by token
	{
		reqStatus := transaction.StatusReq{Reference: req.Reference}
		ret, err := transaction.ApiStatusReq(reqStatus)
		if err != nil {
			//you can retry if occur  Network failure
			fmt.Println(ret, err)
			return
		}

	}
}
