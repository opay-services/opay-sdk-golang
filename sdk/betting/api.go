package betting

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiProviderReq(mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret ProviderResp, err error) {

	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/bills/betting-providers",
		strings.NewReader(""),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", mConf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+mConf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	if logf != nil {
		logf("req", request)
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiBillValidateReq(req BillValidateReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BillVaildateResp, err error) {
	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/bills/validate",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", mConf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+mConf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	if logf != nil {
		logf("req", request)
		logf("req json:", string(jsonReq))
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiBulkBillsReq(req BulkBillsReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BulkStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/bills/bulk-bills",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", mConf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+signStr)
	request.Header.Add("Content-Type", "application/json")

	if logf != nil {
		logf("req", request)
		logf("req json:", string(jsonReq))
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiBulkStatusReq(req OrderStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret BulkStatusResp, err error) {
	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/bills/bulk-status",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", mConf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+mConf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	if logf != nil {
		logf("req", request)
		logf("req json:", string(jsonReq))
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}
