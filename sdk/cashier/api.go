package cashier

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/conf"
	"github.com/opay-services/opay-sdk-golang/sdk"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiCashierInitializeReq(req CashierInitializeReq, opts ...sdk.HttpOption) (ret CashierInitializeResp, err error) {
	if len(req.Currency) == 0 {
		req.Currency = "NGN"
	}

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierStatusReq(req CashierStatusReq, opts ...sdk.HttpOption) (ret CashierStatusResp, err error) {
	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierCloseReq(req CashierCloseReq, opts ...sdk.HttpOption) (ret CashierCloseResp, err error) {
	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierRefundByBankAccountReq(req CashierRefundByBankAccountReq,
	opts ...sdk.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierRefundByOriginReq(req CashierRefundByOriginReq,
	opts ...sdk.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierRefundByOpayMerchantAccountReq(req CashierRefundByOpayMerchantAccountReq,
	opts ...sdk.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierRefundByOpayUserAccountReq(req CashierRefundByOpayUserAccountReq,
	opts ...sdk.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}

	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}

func ApiCashierRefundStatusReq(req CashierRefundStatusReq,
	opts ...sdk.HttpOption) (ret CashierRefundStatusResp, err error) {

	httpClient := sdk.NewHttpClient(opts...)

	jsonReq, err := json.Marshal(&req)
	if err != nil {
		return
	}



	signStr := sdk.Signature(jsonReq)

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

	if conf.GetLog() != nil {
		conf.GetLog()("req", request)
		conf.GetLog()("req json:", string(jsonReq))
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

	if conf.GetLog() != nil {
		conf.GetLog()("ApiCashierRefundStatusReq http response:", string(body))
	}
	err = json.Unmarshal(body, &ret)
	return
}
