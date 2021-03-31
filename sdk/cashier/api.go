package cashier

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiCashierInitializeReq(req CashierInitializeReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierInitializeResp, err error) {
	if len(req.Currency) == 0 {
		req.Currency = "NGN"
	}

	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/initialize",
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

func ApiCashierStatusReq(req CashierStatusReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512(jsonReq, mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/status",
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

func ApiCashierCloseReq(req CashierCloseReq, mConf *conf.OpayMerchantConf, opts ...util.HttpOption) (ret CashierCloseResp, err error) {
	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/close",
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

func ApiCashierRefundByBankAccountReq(req CashierRefundByBankAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOriginReq(req CashierRefundByOriginReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOpayMerchantAccountReq(req CashierRefundByOpayMerchantAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOpayUserAccountReq(req CashierRefundByOpayUserAccountReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundStatusReq(req CashierRefundStatusReq, mConf *conf.OpayMerchantConf,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := mConf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey())

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+"/api/v3/cashier/refund/status",
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
		logf("ApiCashierRefundStatusReq http response:", string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}
