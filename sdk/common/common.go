package common

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)


func PostSignCallOpayGateWay(req interface{}, ret interface{}, mConf *conf.OpayMerchantConf, apiPath string,
	opts ...util.HttpOption) (err error) {

	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	var jsonReq string
	if !reflect.ValueOf(req).IsNil() {
		jsonReq, err = util.OpayJsonMarshal(req)
		if err != nil {
			return err
		}
	}

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+apiPath,
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return err
	}
	request.Header.Add("MerchantId", mConf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+ util.SignatureSHA512([]byte(jsonReq), mConf.GetSecretKey()))
	request.Header.Add("Content-Type", "application/json")

	if logf != nil {
		logf("req", request)
		logf("req json:", string(jsonReq))
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}


	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return nil
}

func PostCallOpayGateWay(req interface{}, ret interface{}, mConf *conf.OpayMerchantConf,
	apiPath string, opts ...util.HttpOption) (err error) {

	logf := mConf.GetLog()
	httpClient := util.NewHttpClient(opts...)

	var jsonReq string
	if req != nil {
		if !reflect.ValueOf(req).IsNil() {
			jsonReq, err = util.OpayJsonMarshal(req)
			if err != nil {
				return err
			}
		}
	}

	request, err := http.NewRequest(
		"POST",
		mConf.GetApiHost()+apiPath,
		strings.NewReader(string(jsonReq)),
	)

	if err != nil {
		return err
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
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if logf != nil {
		logf(string(body))
	}
	err = json.Unmarshal(body, &ret)
	return nil
}
