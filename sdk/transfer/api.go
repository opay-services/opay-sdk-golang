package transfer

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiTransferToOWalletUser(req ToOWalletUserReq, opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {
	httpClient := util.NewHttpClient(opts...)

	logf := conf.GetLog()
	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transfer/toWallet",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	singStr := util.Signature([]byte(jsonReq))
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+singStr)
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

func ApiTransferToOWalletMerchant(req ToOWalletMerchantReq, opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {
	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transfer/toWallet",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	singStr := util.Signature([]byte(jsonReq))
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+singStr)
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

func ApiTransferToBank(req ToBankReq, opts ...util.HttpOption) (ret StatusToBankResp, err error) {
	logf := conf.GetLog()

	if len(req.Receiver.NameCheck) == 0 {
		req.Receiver.NameCheck = "no"
	}

	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}


	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transfer/toBank",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	singStr := util.Signature([]byte(jsonReq))
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + singStr)
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

func ApiStatusToWalletReq(req StatusToWalletReq, opts ...util.HttpOption) (ret StatusToOWalletResp, err error) {
	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transfer/status/toWallet",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature([]byte(jsonReq))
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
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

func ApiStatusToBankReq(req StatusToBankReq, opts ...util.HttpOption) (ret StatusToBankResp, err error) {
	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transfer/status/toBank",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature([]byte(jsonReq))
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
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
