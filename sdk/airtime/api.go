package airtime

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)


func ApiBulkBillsReq(req BulkBillsReq, opts ...util.HttpOption) (ret BulkStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/bills/bulk-bills",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", conf.GetMerchantId())
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

func ApiBulkStatusReq(req OrderStatusReq, opts ...util.HttpOption) (ret BulkStatusResp, err error) {
	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/bills/bulk-status",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+conf.GetPublicKey())
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

