package countries

type CountryInfo struct {
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Name     string `json:"name"`
}

type CountriesSupportResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []CountryInfo `json:"data"`
}
