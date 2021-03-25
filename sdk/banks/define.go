package banks

type BankSupportReq struct {
	CountryCode string `json:"countryCode"`
}

type BankInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type BankSupportResp struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    []BankInfo `json:"data,omitempty"`
}
