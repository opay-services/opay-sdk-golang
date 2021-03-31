package merchant

type InfoMerchantReq struct {
	Email string `json:"email"`
}

type MerchantInfo struct {
	MerchantId   string `json:"merchantId"`
	Email        string `json:"email"`
	BusinessName string `json:"businessName"`
}

type InfoMerchantResp struct {
	Code    string       `json:"code"`
	Data    MerchantInfo `json:"data"`
	Message string       `json:"message"`
}
