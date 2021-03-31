package airtime


var ProviderInfo  = []string{"AIR", "MTN", "GLO", "ETI"}


type BulkDataInfo struct {
	Amount     string `json:"amount"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
	CustomerId string `json:"customerId"`
	Provider   string `json:"provider"`
	Reference  string `json:"reference"`
}

type BulkBillsReq struct {
	CallBackUrl string         `json:"callBackUrl"`
	ServiceType string         `json:"serviceType"`
	BulkData    []BulkDataInfo `json:"bulkData"`
}

type BulkOrderInfo struct {
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
	ErrorMsg  string `json:"errorMsg"`
}

type BulkStatusResp struct {
	Code    string          `json:"code"`
	Data    []BulkOrderInfo `json:"data"`
	Message string          `json:"message"`
}

type OrderStatusReqItem struct {
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference"`
}

type OrderStatusReq struct {
	BulkStatusRequest []OrderStatusReqItem `json:"bulkStatusRequest"`
	ServiceType       string               `json:"serviceType"`
}
