package user

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiUserInfoReq(req InfoUserReq, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/info/user",
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

func ApiUserCreateReq(req UserCreateReq, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/info/user/create",
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

func ApiSendOpt(req SendOptReq, opts ...util.HttpOption) (ret SendOptResp, err error) {

	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/info/user/sendOTP",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+ util.SignatureSHA512([]byte(jsonReq)))
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

func ApiUserUpdateReq(req UserUpdateReq, opts ...util.HttpOption) (ret InfoUserResp, err error) {

	logf := conf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/info/user/update",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+ util.SignatureSHA512([]byte(jsonReq)))
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
