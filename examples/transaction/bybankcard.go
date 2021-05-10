package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/ips"
	"github.com/opay-services/opay-sdk-golang/sdk/transaction"
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

func web()  {
	r := gin.Default()
	r.GET("/dscallback", func(c *gin.Context) {
		buf := make([]byte, 1024)

		n, _ := c.Request.Body.Read(buf)

		fmt.Println(string(buf[0:n]))
	})

	r.POST("/callback", func(c *gin.Context) {
		buf := make([]byte, 1024)

		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))

		fmt.Println(c.Request.Header)
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

func pinopt()  {
	/*bank card transaction with pin + opt
	final status : close, fail, succee
	input status : input-pin, input-phone, input-opt, input-dob, 3DSECURE
	1. create an order
	2. query order status, if order status is pending,
	   You should keep querying the order status
	   until the final status or the required
	   input status (input-pin, input-opt, input-dob, input-phone)
	3.if return input-status, You should notify the user to enter the corresponding code
	  eg. status is "INPUT_PIN", you receive input from your user input, then call ApiInputPin
	  req, then query status, goto step 2

	4.The asynchronous callback address after transaction successful, refer web()/callback
	*/


	//step1: create an order
	req := transaction.ByBankCardReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.Reason = "test"
	req.CustomerEmail = "xxx@163.com"
	req.Return3dsUrl = "https://6f237770df1b.ngrok.io/dscallback"
	req.ExpireAt = "10"
	req.CallbackUrl = "https://6f237770df1b.ngrok.io/callback"
	req.FirstName = "li"
	req.LastName = "jian"
	req.CardCVC = "561"
	req.CardDateYear = "22"
	req.CardDateMonth = "12"
	req.CardNumber = "5061460410121111105"

	ret, err := transaction.ApiByBankCardReq(&req, mConf)
	if err != nil {
		fmt.Println(ret, err)
	}

labstatus:
	//step2: query status util input status or final status
	{
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			reqStatus := transaction.StatusReq{Reference: req.Reference}
			ret, err := transaction.ApiStatusReq(&reqStatus, mConf)
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
				//waiting realized
				return

			case "SUCCESS", "FAIL", "CLOSE":
				return
			}
		}
		//time out
		return
	}

labpin:
	//input pin from user
	{
		fmt.Println("please input pin code from user")
		reqPin := transaction.InputPinReq{Pin: "1105"}
		reqPin.Reference = req.Reference
		ret, err := transaction.ApiInputPinReq(&reqPin, mConf)
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
		//input opt from user
		reqOpt := transaction.InputOtpReq{}
		reqOpt.Reference = req.Reference
		reqOpt.Otp = "543210"
		ret, err := transaction.ApiInputOtpReq(&reqOpt, mConf)
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
		ret, err := transaction.ApiStatusReq(&reqStatus, mConf)
		if err != nil {
			//you can retry if occur  Network failure
			fmt.Println(ret, err)
			return
		}

	}
}

func pin3ds()  {
	/*bank card transaction with pin + opt
	1. create an order
	2. query order status, if order status is pending,
	   You should keep querying the order status
	   until the final status or the required
	   input status (input-pin, input-opt, input-dob, input-phone)
	3.if return input-status, You should notify the user to enter the corresponding code
	  eg. status is "INPUT_PIN", you receive input from your user input, then call ApiInputPin
	  req, then query status, goto step 2

	4.The asynchronous callback address after transaction successful, refer web()/callback
	*/
	req := transaction.ByBankCardReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
	req.Currency = "NGN"
	req.Country = "NG"
	req.Reason = "test"
	req.CustomerEmail = "xxx@163.com"
	req.Return3dsUrl = "https://6f237770df1b.ngrok.io/dscallback"
	req.ExpireAt = "10"
	req.CallbackUrl = "https://6f237770df1b.ngrok.io/callback"
	req.FirstName = "li"
	req.LastName = "jian"
	req.CardCVC = "562"
	req.CardDateYear = "22"
	req.CardDateMonth = "12"
	req.CardNumber = "5061460410121111106"

	ret, err := transaction.ApiByBankCardReq(&req, mConf)
	if err != nil {
		fmt.Println(ret, err)
	}

labstatus:
	//query status
	{
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			reqStatus := transaction.StatusReq{Reference: req.Reference}
			ret, err := transaction.ApiStatusReq(&reqStatus, mConf)
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

			case "3DSECURE":
				goto lab3ds

			case "INPUT-PHONE", "INPUT-DOB", "INPUT-OPT":
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
	//input pin from user
	{

		reqPin := transaction.InputPinReq{Pin: "1106"}
		reqPin.Reference = req.Reference
		ret, err := transaction.ApiInputPinReq(&reqPin, mConf)
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

lab3ds:
	{
		{
			reqStatus := transaction.StatusReq{Reference: req.Reference}
			ret, err := transaction.ApiStatusReq(&reqStatus, mConf)
			if err != nil {
				//you can retry if occur  Network failure
				fmt.Println(ret, err)
				return
			}
			fmt.Println("please press enter...")
			inputReader := bufio.NewReader(os.Stdin)
			input, err := inputReader.ReadString('\n')
			if err != nil{
				fmt.Println(input, err)
			}
		}
	}

labover:
	//query status, Security verification via authurl
	{
		reqStatus := transaction.StatusReq{Reference: req.Reference}
		ret, err := transaction.ApiStatusReq(&reqStatus, mConf)
		if err != nil {
			//you can retry if occur  Network failure
			fmt.Println(ret, err)
			return
		}

	}
}

func main() {
	go web()
	time.Sleep(1*time.Second)
	pinopt()
	fmt.Println("please press enter ...")
	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
}
