package cashier

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiCashierInitializeReq(req CashierInitializeReq, opts ...util.HttpOption) (ret CashierInitializeResp, err error) {
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
		conf.GetApiHost()+"/api/v3/cashier/initialize",
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

func ApiCashierStatusReq(req CashierStatusReq, opts ...util.HttpOption) (ret CashierStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512(jsonReq)

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/status",
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

func ApiCashierCloseReq(req CashierCloseReq, opts ...util.HttpOption) (ret CashierCloseResp, err error) {
	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/close",
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

func ApiCashierRefundByBankAccountReq(req CashierRefundByBankAccountReq,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOriginReq(req CashierRefundByOriginReq,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOpayMerchantAccountReq(req CashierRefundByOpayMerchantAccountReq,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundByOpayUserAccountReq(req CashierRefundByOpayUserAccountReq,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/refund",
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

func ApiCashierRefundStatusReq(req CashierRefundStatusReq,
	opts ...util.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := util.NewHttpClient(opts...)
	logf := conf.GetLog()

	jsonReq, err := util.OpayJsonMarshal(&req)
	if err != nil {
		return
	}

	signStr := util.SignatureSHA512([]byte(jsonReq))

	request, err := http.NewRequest(
		"POST",
		conf.GetApiHost()+"/api/v3/cashier/refund/status",
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
		logf("ApiCashierRefundStatusReq http response:", string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}
