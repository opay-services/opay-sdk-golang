package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opay-services/opay-sdk-golang/sdk/airtime"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/ips"
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
			if notify.VerfiySignature(){
				//TODO
			}
		}
	})
	r.Run(":8080")
}

func main()  {
	go web()
	req := airtime.BulkBillsReq{}
	req.ServiceType = "airtime"
	req.CallBackUrl = "https://6f237770df1b.ngrok.io/callback"
	req.BulkData = make([]airtime.BulkDataInfo, 2, 2)

	reqStatus := airtime.OrderStatusReq{}
	reqStatus.ServiceType = "airtime"
	reqStatus.BulkStatusRequest = make([]airtime.OrderStatusReqItem, 2, 2)

	for i := 0; i < 2; i++ {
		req.BulkData[i].Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano()+int64(i))
		reqStatus.BulkStatusRequest[i].Reference = req.BulkData[i].Reference
		req.BulkData[i].Provider = "AIR"
		req.BulkData[i].Amount = "50000"
		req.BulkData[i].Country = "NG"
		req.BulkData[i].Currency = "NGN"
		req.BulkData[i].CustomerId = "20019212912901281821982" + fmt.Sprintf("%v", i)
	}

	ret, err := airtime.ApiBulkBillsReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}

	fmt.Println("please press enter...")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Println(input, err)
	}

	{
		ret, err := airtime.ApiBulkStatusReq(reqStatus)
		if err != nil {
			fmt.Println(ret, err)
		}
	}
}
