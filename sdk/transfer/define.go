package transfer

type OWalletReceiverUser struct {
	/*Receiver's name, utf-8 encoded*/
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
}

type ToOWalletUserReq struct {
	/*Amount to transfer in kobo*/
	Amount   string              `json:"amount"`
	Country  string              `json:"country"`
	Currency string              `json:"currency"`
	Reason   string              `json:"reason,omitempty"`
	Receiver OWalletReceiverUser `json:"receiver"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference"`
}

type OWalletReceiverMerchant struct {
	/*Needed while Receiver's role is Merchant,
	Receiver's OPay merchantId which could enquiry via /info/merchant
	endpoint. e.g. 256619092316009
	*/
	MerchantId string `json:"merchantId"`
	/*Receiver's name, utf-8 encoded*/
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type ToOWalletMerchantReq struct {
	/*Amount to transfer in kobo*/
	Amount   string                  `json:"amount"`
	Country  string                  `json:"country"`
	Currency string                  `json:"currency"`
	Reason   string                  `json:"reason,omitempty"`
	Receiver OWalletReceiverMerchant `json:"receiver"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference"`
}

type StatusToOWalletData struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	FailureReason string `json:"failureReason,omitempty"`
	Fee           string `json:"fee"`
	MerchantId    string `json:"merchantId,omitempty"`
	OrderNo       string `json:"orderNo"`
	PhoneNumber   string `json:"phoneNumber,omitempty"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	Type          string `json:"type"`
}

type StatusToOWalletResp struct {
	Code string `json:"code"`
	Data StatusToOWalletData `json:"data"`
	Message string `json:"message"`
}

type BankReceiver struct {
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	/*Receiver's name, utf-8 encoded*/
	Name string `json:"name,omitempty"`
	/*If check name
	  no
	  yes
	  fuzzy
	  Default is no
	*/
	NameCheck string `json:"nameCheck"`
}

type ToBankReq struct {
	/*Amount to transfer in kobo*/
	Amount   string       `json:"amount"`
	Country  string       `json:"country"`
	Currency string       `json:"currency"`
	Reason   string       `json:"reason,omitempty"`
	Receiver BankReceiver `json:"receiver"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference"`
}

type StatusToBankData struct {
	Amount            string `json:"amount"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	Currency          string `json:"currency"`
	FailureReason     string `json:"failureReason,omitempty"`
	Fee               string `json:"fee"`
	OrderNo           string `json:"orderNo"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

type StatusToBankResp struct {
	Code string `json:"code"`
	Data StatusToBankData `json:"data"`
	Message string `json:"message"`
}

type StatusToWalletReq struct {
	OrderNo   string `json:"orderNo,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type StatusToBankReq struct {
	OrderNo   string `json:"orderNo,omitempty"`
	Reference string `json:"reference,omitempty"`
}
