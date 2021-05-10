package user

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiUserInfoReq(req *InfoUserReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	err = common.PostCallOpayGateWay(req, &ret, mConf, "/api/v3/info/user", opts...)
	return
}

func ApiUserCreateReq(req *UserCreateReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	err = common.PostCallOpayGateWay(req, &ret, mConf, "/api/v3/info/user/create", opts...)
	return
}

func ApiSendOpt(req *SendOptReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret SendOptResp, err error) {


	err = common.PostSignCallOpayGateWay(req, &ret, mConf, "/api/v3/info/user/sendOTP", opts...)
	return
}

func ApiUserUpdateReq(req *UserUpdateReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	err = common.PostSignCallOpayGateWay(req, &ret, mConf, "/api/v3/info/user/update", opts...)
	return
}
