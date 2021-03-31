package charge

type InitializeReq struct {
	Reference     string `json:"reference"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	UserRequestIp string `json:"userRequestIp"`
	CallbackUrl   string `json:"callbackUrl"`
	ReturnUrl     string `json:"returnUrl"`
	ExpireAt      string `json:"expireAt"`
	ChargerType   string `json:"chargerType"`
	ChargerId     string `json:"chargerId"`
}

type InitializeInfo struct {
	Reference     string `json:"reference"`
	OrderNo       string `json:"orderNo"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	FailureReason string `json:"failureReason"`
	CashierUrl    string `json:"cashierUrl"`
	chargerType   string `json:"chargerType"`
}

type InitializeResp struct {
	Code    string         `json:"code"`
	Data    InitializeInfo `json:"data"`
	Message string         `json:"message"`
}

type StatusReq struct {
	Reference *string  `json:"reference,omitempty"`
	OrderNo   *string `json:"orderNo,omitempty"`
}

type StatusInfo struct {
	Reference     string `json:"reference"`
	OrderNo       string `json:"orderNo"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	FailureReason string `json:"failureReason"`
}

type StatusResp struct {
	Code    string     `json:"code"`
	Data    StatusInfo `json:"data"`
	Message string     `json:"message"`
}
