package transaction

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiByBankCardReq(req ByBankCardReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret TranscationChecktoutResp, err error) {

	if req.PayType != "bankcard" {
		req.PayType = "bankcard"
	}
	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/initialize", opts...)
	return
}

func ApiByBankAccountReq(req ByBankAccountReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret TranscationChecktoutResp, err error) {

	if req.PayType != "bankaccount" {
		req.PayType = "bankaccount"
	}
	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/initialize", opts...)
	return
}

func ApiInputOtpReq(req InputOtpReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InputOptResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/input-otp", opts...)
	return
}

func ApiInputPhoneReq(req InputPhoneReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InputPhoneResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/input-phone", opts...)
	return
}

func ApiInputPinReq(req InputPinReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InputPinResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/input-pin", opts...)
	return
}

func ApiInputDobReq(req InputDobReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InputDobResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/input-dob", opts...)
	return
}

func ApiStatusReq(req StatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret StatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/status", opts...)
	return
}

func ApiSupportBanksReq(req SupportBanksReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret SupportBankResp, err error) {

	err = common.PostCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/banks", opts...)
	return
}

func ApiUssdCodeReq(req UssdCodeReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret UssdCodeResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/ussd/initialize", opts...)
	return
}

func ApiUssdOrderStatusReq(req UssdOrderStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret UssdCodeResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/ussd/status", opts...)
	return
}

func ApiBankTransferInitializeReq(req BankTransferInitializeReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (
	ret BankTransferInitializeResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/bankTransfer/initialize", opts...)
	return
}

func ApiBankTransferStatusReq(req BankTransferStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (
	ret BankTransferStatusResp, err error) {

	err = common.PostSignCallOpayGateWay(&req, &ret, mConf, "/api/v3/transaction/bankTransfer/status", opts...)
	return
}
