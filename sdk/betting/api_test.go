package betting

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
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

func TestApiProviderReq(t *testing.T) {
	ret, err := ApiProviderReq()
	if err != nil {
		fmt.Println(ret, err)
	}
}

func TestApiBillValidateReq(t *testing.T) {
	req := BillValidateReq{}
	req.Provider = "NAIRABET"
	req.ServiceType = "betting"
	req.CustomerId = "20019212912901281821982"

	ret, err := ApiBillValidateReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}

func TestApiBettingBulkBillsReq(t *testing.T) {
	req := BulkBillsReq{}
	req.ServiceType = "betting"
	req.CallBackUrl = "http://localhost:8000"
	req.BulkData = make([]BulkDataInfo, 2, 2)
	for i := 0; i < 2; i++ {
		req.BulkData[i].Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano()+int64(i))
		req.BulkData[i].Provider = "NAIRABET"
		req.BulkData[i].Amount = "50000"
		req.BulkData[i].Country = "NG"
		req.BulkData[i].Currency = "NGN"
		req.BulkData[i].CustomerId = "20019212912901281821982" + fmt.Sprintf("%v", i)
	}

	ret, err := ApiBulkBillsReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}

func TestApiAirtimeTopup(t *testing.T)  {
	req := BulkBillsReq{}
	req.ServiceType = "airtime"
	req.CallBackUrl = "http://localhost:8000"
	req.BulkData = make([]BulkDataInfo, 2, 2)
	for i := 0; i < 2; i++ {
		req.BulkData[i].Reference = fmt.Sprintf("testlijian_%v", time.Now().UnixNano()+int64(i))
		req.BulkData[i].Provider = "AIR"
		req.BulkData[i].Amount = "50000"
		req.BulkData[i].Country = "NG"
		req.BulkData[i].Currency = "NGN"
		req.BulkData[i].CustomerId = "20019212912901281821982" + fmt.Sprintf("%v", i)
	}

	ret, err := ApiBulkBillsReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}

func TestApiBulkStatusReq(t *testing.T) {
	req := OrderStatusReq{}
	req.ServiceType = "betting"
	req.BulkStatusRequest = make([]OrderStatusReqItem, 2, 2)

	req.BulkStatusRequest[0].Reference = "testlijian_1617106425608068000"
	req.BulkStatusRequest[1].Reference = "testlijian_1617106425608075001"

	ret, err := ApiBulkStatusReq(req)
	if err != nil {
		fmt.Println(ret, err)
	}
}
