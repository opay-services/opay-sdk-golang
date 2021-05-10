package charge

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiInitializeReq(req *InitializeReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InitializeResp, err error) {

	err = common.PostCallOpayGateWay(req, &ret, mConf, "/api/v3/charge/initialize", opts...)
	return
}

func ApiStatusReq(req *StatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret StatusResp, err error) {

	err = common.PostSignCallOpayGateWay(req, &ret, mConf, "/api/v3/charge/status", opts...)
	return
}

