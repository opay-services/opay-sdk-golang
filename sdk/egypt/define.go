package egypt

type EgyptOrderStatusReq struct {
	Reference *string `json:"reference,omitempty"`
	OrderNo   *string `json:"orderNo,omitempty"`
}

type EgyptOrderStatusInfo struct {
	MerchantId string `json:"merchantId"`
	Reference  string `json:"reference"`
	OrderNo    string `json:"orderNo"`
	Status     string `json:"status"`
	AuthUrl    string `json:"authUrl"`
}

type EgtypOrderStatusResp struct {
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Data    *EgyptOrderStatusInfo `json:"data,omitempty"`
}

type EgyptAmountStruct struct {
	Currency string `json:"currency"`
	Total    uint   `json:"total"`
}
