package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opay-services/opay-sdk-golang/sdk/betting"
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
	req := betting.BulkBillsReq{}
	req.ServiceType = "betting"
	req.CallBackUrl = "https://6f237770df1b.ngrok.io/callback"
	req.BulkData = make([]betting.BulkDataInfo, 2, 2)

	reqStatus := betting.OrderStatusReq{}
	reqStatus.ServiceType = "betting"
	reqStatus.BulkStatusRequest = make([]betting.OrderStatusReqItem, 2, 2)

	for i := 0; i < 2; i++ {
		req.BulkData[i].Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano()+int64(i))
		reqStatus.BulkStatusRequest[i].Reference = req.BulkData[i].Reference
		req.BulkData[i].Provider = "NAIRABET"
		req.BulkData[i].Amount = "50000"
		req.BulkData[i].Country = "NG"
		req.BulkData[i].Currency = "NGN"
		req.BulkData[i].CustomerId = "20019212912901281821982" + fmt.Sprintf("%v", i)
	}

	ret, err := betting.ApiBulkBillsReq(&req, mConf)
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
		ret, err := betting.ApiBulkStatusReq(&reqStatus, mConf)
		if err != nil {
			fmt.Println(ret, err)
		}
	}
}
