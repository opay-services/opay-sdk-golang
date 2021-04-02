package banks

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiGetBankList(countryCode string, mconf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BankSupportResp, err error){
	req := BankSupportReq{CountryCode:countryCode}
	err = common.PostCallOpayGateWay(&req, &ret, mconf, "/api/v3/banks", opts...)
	return
}
