package betting

type BettingProviderInfo struct {
	/*Betting topup service provider*/
	Provider string `json:"provider"`
	/*Provider logo URL*/
	ProviderLogoUrl string `json:"providerLogoUrl"`
	/*amount in kobo. The minimum limit for a single transaction, if it is not returned, it means there is no limit*/
	MinAmount string `json:"minAmount"`
	/*amount in kobo. The maximum limit for a single transaction, if it is not returned, it means there is no limit*/
	MaxAmount string `json:"maxAmount"`
}

type BettingProviderResp struct {
	Code    string                `json:"code"`
	Data    []BettingProviderInfo `json:"data"`
	Message string                `json:"message"`
}

type BillValidateReq struct {
	/*betting*/
	ServiceType string `json:"serviceType"`
	/*provider returned in betting-providers*/
	Provider string `json:"provider"`
	/*benificiency's customerId in specific provider*/
	CustomerId string `json:"customerId"`
}

type BillVaildateInfo struct {
	Provider   string `json:"provider"`
	CustomerId string `json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	UserName   string `json:"userName"`
}

type BillVaildateResp struct {
	Code    string           `json:"code"`
	Data    BillVaildateInfo `json:"data"`
	Message string           `json:"message"`
}

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
