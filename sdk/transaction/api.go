package transaction

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiByBankCardReq(req ByBankCardReq, opts ...util.HttpOption) (ret TranscationChecktoutResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	if req.PayType != "bankcard"{
		req.PayType = "bankcard"
	}

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/initialize",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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


func ApiByBankAccountReq(req ByBankAccountReq, opts ...util.HttpOption)(ret TranscationChecktoutResp, err error)  {

	httpClient := util.NewHttpClient(opts...)
	if req.PayType != "bankaccount"{
		req.PayType = "bankaccount"
	}

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/initialize",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiInputOtpReq(req InputOtpReq, opts ...util.HttpOption) (ret InputOptResp, err error){

	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/input-otp",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiInputPhoneReq(req InputPhoneReq, opts ...util.HttpOption) (ret InputPhoneResp, err error) {

	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/input-phone",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiInputPinReq(req InputPinReq, opts ...util.HttpOption) (ret InputPinResp, err error) {

	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/input-pin",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiInputDobReq(req InputDobReq, opts ...util.HttpOption) (ret InputDobResp, err error) {

	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/input-dob",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiStatusReq(req StatusReq, opts ...util.HttpOption)(ret StatusResp, err error)  {

	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/status",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature(jsonReq)

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiSupportBanksReq(req SupportBanksReq, opts ...util.HttpOption) (ret SupportBankResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/banks",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiUssdCodeReq(req UssdCodeReq, opts ...util.HttpOption) (ret UssdCodeResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/ussd/initialize",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature(jsonReq)

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiUssdOrderStatusReq(req UssdOrderStatusReq, opts ...util.HttpOption) (ret UssdCodeResp, err error) {
	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/ussd/status",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature(jsonReq)

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiBankTransferInitializeReq(req BankTransferInitializeReq, opts ...util.HttpOption) (
	ret BankTransferInitializeResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/bankTransfer/initialize",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature(jsonReq)
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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

func ApiBankTransferStatusReq(req BankTransferStatusReq, opts ...util.HttpOption) (
	ret BankTransferStatusResp, err error) {


	httpClient := util.NewHttpClient(opts...)
	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/transaction/bankTransfer/status",
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return
	}

	signStr := util.Signature(jsonReq)

	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer " + signStr)
	request.Header.Add("Content-Type", "application/json")

	logf := conf.GetLog()
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