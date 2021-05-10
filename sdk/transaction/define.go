package transaction

type ByBankCardReq struct {
	Amount   string `json:"amount"`
	BankCode string `json:"bankCode"`
	/*The house/building address registered with the card.*/
	BillingAddress string `json:"billingAddress,omitempty"`
	/*The city registered with the card.*/
	BillingCity string `json:"billingCity,omitempty"`
	/*The country registered with the card. e.g. US, CA, AU, etc.*/
	BillingCountry string `json:"billingCountry,omitempty"`
	/*The state registered with the card. e.g. NY for New York, CA for California, etc.*/
	BillingState string `json:"billingState,omitempty"`
	/*Zip code or postal card registered with the card.*/
	BillingZip string `json:"billingZip,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	/*The asynchronous callback address after transaction successful.*/
	CallbackUrl string `json:"callbackUrl"`
	/*Bankcard expired month e.g. 01*/
	CardCVC       string `json:"cardCVC"`
	CardDateMonth string `json:"cardDateMonth"`
	/*Bankcard expired year e.g. 22*/
	CardDateYear  string `json:"cardDateYear"`
	CardNumber    string `json:"cardNumber"`
	Country       string `json:"country"`
	Currency      string `json:"currency"`
	CustomerEmail string `json:"customerEmail,omitempty"`
	CustomerPhone string `json:"customerPhone,omitempty"`
	ExpireAt      string `json:"expireAt"`
	PayType       string `json:"payType"`
	Reason        string `json:"reason"`
	Reference     string `json:"reference"`
	/*What page do you need to call back to after 3ds verification is successful*/
	Return3dsUrl string `json:"return3dsUrl,omitempty"`
}

type ByBankAccountReq struct {
	/*Amount to transfer in kobo*/
	Amount            string `json:"amount"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	Bvn               string `json:"bvn"`
	CallbackUrl       string `json:"callbackUrl"`
	Country           string `json:"country"`
	Currency          string `json:"currency"`
	CustomerPhone     string `json:"customerPhone"`
	DobDay            string `json:"dobDay"`
	DobMonth          string `json:"dobMonth"`
	DobYear           string `json:"dobYear"`
	ExpireAt          string `json:"expireAt"`
	PayType           string `json:"payType"`
	Reason            string `json:"reason"`
	Reference         string `json:"reference"`
	Return3dsUrl      string `json:"return3dsUrl,omitempty"`
}

type CheckoutInfo struct {
	/*Amount to transfer in kobo*/
	Amount        string `json:"amount"`
	AuthURL       string `json:"authURL"`
	Currency      string `json:"currency"`
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when tokenize is true*/
	Token string `json:"token"`
}

type TranscationChecktoutResp struct {
	Code    string        `json:"code"`
	Data    *CheckoutInfo `json:"data,omitempty"`
	Message string        `json:"message"`
}

type InputOtpReq struct {
	OrderNo   string `json:"orderNo"`
	Otp       string `json:"otp"`
	Reference string `json:"reference,omitempty"`
}

type OptInfo struct {
	authURL       string `json:"authURL"`
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when status is SUCCESSFUL and tokenize is true*/
	Token string `json:"token"`
}

type InputOptResp struct {
	Code    string   `json:"code"`
	Data    *OptInfo `json:"data,omitempty"`
	Message string   `json:"message"`
}

type InputPhoneReq struct {
	OrderNo   string `json:"orderNo"`
	Phone     string `json:"phone"`
	Reference string `json:"reference,omitempty"`
}

type InputPhoneInfo struct {
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when status is SUCCESSFUL and tokenize is true*/
	Token string `json:"token"`
}

type InputPhoneResp struct {
	Code    string         `json:"code"`
	Data    InputPhoneInfo `json:"data,omitempty"`
	Message string         `json:"message"`
}

type InputPinReq struct {
	OrderNo   string `json:"orderNo"`
	Pin       string `json:"pin"`
	Reference string `json:"reference,omitempty"`
}

type PinInfo struct {
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when status is SUCCESSFUL and tokenize is true*/
	Token string `json:"token"`
}

type InputPinResp struct {
	Code    string   `json:"code"`
	Data    *PinInfo `json:"data,omitempty"`
	Message string   `json:"message"`
}

type InputDobReq struct {
	Dob       string `json:"dob"`
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference,omitempty"`
}

type DobInfo struct {
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when status is SUCCESSFUL and tokenize is true*/
	Token string `json:"token"`
}
type InputDobResp struct {
	Code    string   `json:"code"`
	Data    *DobInfo `json:"data"`
	Message string   `json:"message"`
}

type StatusReq struct {
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference,omitempty"`
}

type StatusInfo struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	FailureReason string `json:"failureReason"`
	OrderNo       string `json:"orderNo"`
	Reference     string `json:"reference"`
	/*
		INITIAL
		PENDING
		SUCCESS
		FAIL
		INPUT-PIN
		INPUT-OTP
		INPUT-PHONE
		INPUT-DOB
		3DSECURE
		CLOSE
	*/
	Status string `json:"status"`
	/*Only returned when status is SUCCESSFUL and tokenize is true*/
	Token string `json:"token"`
}

type StatusResp struct {
	Code    string      `json:"code"`
	Data    *StatusInfo `json:"data"`
	Message string      `json:"message"`
}

type SupportBanksReq struct {
	Country string `json:"country"`
}

type BankInfo struct {
	Code  string `json:"code"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type BankInfos struct {
	Banks []BankInfo `json:"banks"`
}

type SupportBankResp struct {
	Code    string    `json:"code"`
	Data    BankInfos `json:"data"`
	Message string    `json:"message"`
}

type UssdCodeReq struct {
	Amount        string `json:"amount"`
	BankCode      string `json:"bankCode"`
	CallbackUrl   string `json:"callbackUrl"`
	Currency      string `json:"currency"`
	ExpireAt      string `json:"expireAt"`
	ProductDesc   string `json:"productDesc"`
	Reference     string `json:"reference"`
	UserPhone     string `json:"userPhone"`
	UserRequestIp string `json:"userRequestIp"`
}

type UssdCodeInfo struct {
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	OrderNo      string `json:"orderNo"`
	Reference    string `json:"reference"`
	Status       string `json:"status"`
	UssdDialCode string `json:"ussdDialCode"`
}

type UssdCodeResp struct {
	Code    string        `json:"code"`
	Data    *UssdCodeInfo `json:"data,omitempty"`
	Message string        `json:"message"`
}

type UssdOrderStatusReq struct {
	OrderNo   string `json:"orderNo,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type UssdOrderStatusInfo struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

type UssdOrderStatusResp struct {
	Code    string               `json:"code"`
	Data    *UssdOrderStatusInfo `json:"data,omitempty"`
	Message string               `json:"message"`
}

type BankTransferInitializeReq struct {
	Amount      string `json:"amount"`
	CallbackUrl string `json:"callbackUrl"`
	Currency    string `json:"currency"`
	ExpireAt    string `json:"expireAt"`
	ProductDesc string `json:"productDesc"`
	Reference   string `json:"reference"`
	/*User phone number sent by merchant*/
	UserPhone string `json:"userPhone"`
	/*The IP address requested by user, need pass-through by merchant, user Anti-phishing verification.*/
	UserRequestIp string `json:"userRequestIp"`
}

type BankTransferInitializeInfo struct {
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	OrderNo         string `json:"orderNo"`
	Reference       string `json:"reference"`
	Status          string `json:"status"`
	TransferAccount string `json:"transferAccount"`
	TransferBank    string `json:"transferBank"`
}

type BankTransferInitializeResp struct {
	Code    string                      `json:"code"`
	Data    *BankTransferInitializeInfo `json:"data,omitempty"`
	Message string                      `json:"message"`
}

type BankTransferStatusReq struct {
	OrderNo   string `json:"orderNo,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type BankTransferOrderStatusInfo struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	OrderNo   string `json:"orderNo"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

type BankTransferStatusResp struct {
	Code    string                       `json:"code"`
	Data    *BankTransferOrderStatusInfo `json:"data,omitempty"`
	Message string                       `json:"message"`
}

type EgyptAmountStruct struct {
	Currency string `json:"currency"`
	Total    uint   `json:"total"`
}

type EgyptProductStruct struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type EgyptTransactionCreateReq struct {
	ReturnUrl    string              `json:"returnUrl,omitempty"`
	CallbackUrl  string              `json:"callbackUrl,omitempty"`
	Amount       *EgyptAmountStruct  `json:"amount"`
	Product      *EgyptProductStruct `json:"product,omitempty"`
	Remark       *string             `json:"remark,omitempty"`
	UserClientIP *string             `json:"userClientIP,omitempty"`
	Reference    string              `json:"reference"`
}

type EgyptOrderStatusReq struct {
	Reference string `json:"reference"`
	OrderNo   string `json:"orderNo"`
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
