package transfer

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiTransferToOWalletUser(req ToOWalletUserReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transfer/toWallet", opts...)
	return
}

func ApiTransferToOWalletMerchant(req ToOWalletMerchantReq, mConf *conf.OpayMerchantConf,opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transfer/toWallet", opts...)
	return
}

func ApiTransferToBank(req ToBankReq, mConf *conf.OpayMerchantConf,opts ...util.HttpOption) (ret StatusToBankResp, err error) {
	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transfer/toBank", opts...)
	return
}

func ApiStatusToWalletReq(req StatusToWalletReq, mConf *conf.OpayMerchantConf,opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {
	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transfer/status/toWallet", opts...)
	return
}

func ApiStatusToBankReq(req StatusToBankReq, mConf *conf.OpayMerchantConf,opts ...util.HttpOption) (ret StatusToBankResp, err error) {
	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transfer/status/toBank", opts...)
	return
}
