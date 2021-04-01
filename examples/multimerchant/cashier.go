package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opay-services/opay-sdk-golang/sdk/cashier"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/ips"
	"math/rand"
	"os"
	"time"
)

/*
* note: 1.same merchantId only init once
        2.same currency get will fetch latest conf
 */
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
		"USD",
	)

	m2.SetLog(func(a ...interface{}) {
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

		notify := ips.MerchantAcquiring{}
		err := json.Unmarshal(buf[:n], &notify)
		if err != nil {
			fmt.Println(err)
		}else {
			/*
				 if you support multi currency and any currency merchant config only one,
				   you can call conf.GetMerchantConfigByCurrency, others, you should
					call conf.GetMerchantConfigByMerchantId。this means, you must know
					Corresponding merchant id
			*/

			//if same currency only one
			{
				mConf := conf.GetMerchantConfigByCurrency(notify.Payload.Currency)
				notify.VerfiySignature(mConf)
			} //else{
				/*
					//get merchant id by reference, you self define fn
					mConf := conf.GetMerchantConfigByMerchantId()
					notify.VerfiySignature(mConf)
				 */
			//}
		}
	})
	r.Run(":8080")
}

func main() {
	go web()
	m1 := conf.GetMerchantConfigByMerchantId("256620112018025")
	m2 := conf.GetMerchantConfigByMerchantId("256621033118750")
	goto lab2

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

	}

	lab2:{
		{
			//create cashier order
			req := cashier.CashierInitializeReq{}
			req.Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano())
			req.MchShortName = "lijian_test"
			req.ProductName = "goods"
			req.ProductDesc = "test goods"
			req.UserPhone = "+2349876543210"
			req.UserRequestIp = "123.123.123.123"
			req.Amount = "8"
			req.Currency = "USD"
			req.PayTypes = []string{"BalancePayment", "BonusPayment", "OWealth"}
			req.PayMethods = []string{"account", "qrcode", "bankCard", "bankAccount"}
			req.ExpireAt = "10"
			req.CallbackUrl = "https://6f237770df1b.ngrok.io/callback"
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

			fmt.Println("please press enter ...")
			inputReader := bufio.NewReader(os.Stdin)
			inputReader.ReadString('\n')
		}
	}
}
