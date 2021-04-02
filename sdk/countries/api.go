package countries

import (
	"github.com/opay-services/opay-sdk-golang/sdk/common"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

func ApiGetContriesSupport(mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CountriesSupportResp, err error) {

	err = common.PostCallOpayGateWay(nil, &ret, mConf, "/api/v3/countries", opts...)
	return
}
