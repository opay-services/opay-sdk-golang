package airtime

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)


func ApiBulkBillsReq(req *BulkBillsReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BulkStatusResp, err error) {
	err = common.PostSignCallOpayGateWay(req, &ret, mConf, "/api/v3/bills/bulk-bills", opts...)
	return
}

func ApiBulkStatusReq(req *OrderStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BulkStatusResp, err error) {

	err = common.PostCallOpayGateWay(req, &ret, mConf, "/api/v3/bills/bulk-status", opts...)
	return
}

