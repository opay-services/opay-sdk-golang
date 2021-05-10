package user

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
	"time"
)

var mConf *conf.OpayMerchantConf
func init()  {
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

func TestApiUserInfoReq(t *testing.T) {
	req := InfoUserReq{}
	req.PhoneNumber = "+2349876543210"
	resp, err := ApiUserInfoReq(&req, mConf)
	fmt.Println(resp, err)

	if err != nil {
		return
	}

	if resp.Code == "user_not_found" {

	}
}

func TestApiSendOpt(t *testing.T) {
	req := SendOptReq{}
	req.PhoneNumber = "+2349876543210"

	resp, err := ApiSendOpt(&req, mConf)
	fmt.Println(resp, err)
}

func TestApiUserCreateReq(t *testing.T) {
	req := UserCreateReq{}
	req.PhoneNumber = "+2349876543865"
	req.Otp = "123456"
	req.FirstName = "li"
	req.LastName = "jian"
	req.Email = "lijian2233@163.com"
	req.Address = "xxx"
	req.Password = "123456"
	resp, err := ApiUserCreateReq(&req, mConf)
	fmt.Println(resp, err)
}

func TestApiUserUpdateReq(t *testing.T) {
	req := UserUpdateReq{}
	req.PhoneNumber = "+2349876543865"
	req.FirstName = "li"
	req.LastName = "jian"
	req.Email = "lijian2233@163.com"
	req.Address = "xxx"
	resp, err := ApiUserUpdateReq(&req, mConf)
	fmt.Println(resp, err)

}
