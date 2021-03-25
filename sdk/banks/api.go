package banks

import (
	"encoding/json"
	"github.com/opay-services/opay-sdk-golang/conf"
	"github.com/opay-services/opay-sdk-golang/sdk"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiGetBankList(countryCode string, opts ...sdk.HttpOption) (ret BankSupportResp, err error){
	httpClient := sdk.NewHttpClient(opts...)

	req := BankSupportReq{CountryCode:countryCode}

	jsonReq, err := json.Marshal(&req)
	if err != nil{
		return
	}

	request, err := http.NewRequest("POST",
		conf.GetApiHost()+"/api/v3/banks",
		strings.NewReader(string(jsonReq)))

	if err != nil{
		return
	}
	request.Header.Add("MerchantId", conf.GetMerchantId())
	request.Header.Add("Authorization", "Bearer "+conf.GetPublicKey())
	request.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &ret)
	return
}