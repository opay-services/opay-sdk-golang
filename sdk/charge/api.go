package charge

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiInitializeReq(req InitializeReq, opts ...util.HttpOption) (ret InitializeResp, err error) {
	if len(req.Currency) == 0 {
		req.Currency = "NGN"
	}

	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/charge/initialize",
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

func ApiStatusReq(req StatusReq, opts ...util.HttpOption) (ret StatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512(jsonReq)

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/charge/status",
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

