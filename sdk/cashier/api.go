package cashier

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiCashierInitializeReq(req CashierInitializeReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierInitializeResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/initialize")
	return
}


func ApiCashierStatusReq(req CashierStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/status", opts...)
	return
}

func ApiCashierCloseReq(req CashierCloseReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierCloseResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/close", opts...)
	return
}

func ApiCashierRefundByBankAccountReq(req CashierRefundByBankAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/refund", opts...)
	return
}

func ApiCashierRefundByOriginReq(req CashierRefundByOriginReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/refund", opts...)
	return
}

func ApiCashierRefundByOpayMerchantAccountReq(req CashierRefundByOpayMerchantAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/refund", opts...)
	return
}

func ApiCashierRefundByOpayUserAccountReq(req CashierRefundByOpayUserAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/refund", opts...)
	return
}

func ApiCashierRefundStatusReq(req CashierRefundStatusReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/cashier/refund/status", opts...)
	return
}
