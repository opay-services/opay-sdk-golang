package transfer

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

func TestApiTransferToOWalletUser(t *testing.T) {
	req := ToOWalletUserReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = OWalletReceiverUser{
		PhoneNumber:"+2348160564736",
		Type:"USER",
		Name:"Andy Lee",
	}
	ApiTransferToOWalletUser(req)
}

func TestApiTransferToOWalletMerchant(t *testing.T) {
	req := ToOWalletMerchantReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = OWalletReceiverMerchant{
		MerchantId:"256620112018024",
		Type:"MERCHANT",
		Name:"Andy Lee",
	}
	ApiTransferToOWalletMerchant(req)
}

func TestApiTransferToBank(t *testing.T) {
	req := ToBankReq{}
	req.Amount = "100"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Currency = "NGN"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = BankReceiver{
		BankCode:"033",
		BankAccountNumber:"2070043686",
		NameCheck:"yes",
		Name:"test",
	}

	rsp, _ := ApiTransferToBank(req)
	fmt.Println(rsp)
}

func TestApiStatusToWalletReq(t *testing.T) {
	req := StatusToWalletReq{Reference:"testlijian_1616665358658881000"}
	ApiStatusToWalletReq(req)
}

func TestApiStatusToBankReq(t *testing.T) {
	req := StatusToBankReq{Reference:"testlijian_1616665147758091000"}
	ApiStatusToBankReq(req)
}