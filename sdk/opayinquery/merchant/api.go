package merchant

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiInfoMerchantReq(req *InfoMerchantReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InfoMerchantResp, err error)   {

	err = common.PostCallOpayGateWay(req, &ret, mConf, "/api/v3/info/merchant", opts...)
	return
}
