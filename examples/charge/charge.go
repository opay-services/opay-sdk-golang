package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opay-services/opay-sdk-golang/sdk/charge"
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

func web()  {
	r := gin.Default()

	r.POST("/callback", func(c *gin.Context) {
		buf := make([]byte, 1024)

		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))

		notify := ips.BettingAndAirTime{}
		err := json.Unmarshal(buf[:n], &notify)
		if err != nil {
			fmt.Println(err)
		}else {
			if notify.VerfiySignature(mConf){
				//TODO
			}
		}
	})
	r.Run(":8080")
}

func main()  {

	go web()

	req := charge.InitializeReq{}
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.UserRequestIp = "123.123.123.123"
	req.Amount = "100"
	req.Currency = "NGN"
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.ReturnUrl = "http://localhost:8080"
	req.ChargerType = "USER"
	req.ChargerId = "156619102400201625"

	ret, err := charge.ApiInitializeReq(req, mConf)

	if err != nil {
		fmt.Println(ret, err)
	}

	if ret.Code != "00000"{
		return
	}

	fmt.Println("please press enter...")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Println(input, err)
	}

	{
		reqStatus := charge.StatusReq{}
		reqStatus.Reference = &req.Reference
		resp, err := charge.ApiStatusReq(reqStatus, mConf)
		if err != nil{
			fmt.Println(err)
			return
		}

		if resp.Code != "00000"{
			fmt.Println(resp.Message)
		}
	}
}
