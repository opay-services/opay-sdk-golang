package cashier

/*do not modity field order, json marshal struct by field order*/
type CashierInitializeReq struct {
	/*Amount in kobo*/
	Amount string `json:"amount,omitempty"`
	/*The asynchronous callback address after transaction successful.*/
	CallbackUrl string `json:"callbackUrl,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Transaction would be closed within specific time. Value is in minute.*/
	ExpireAt string `json:"expireAt,omitempty"`
	/*The short name of a Merchant. It's displayed on the payment confirmation page.*/
	MchShortName string `json:"mchShortName,omitempty"`

	/*Payment method
	  account (Balance payment)
	  qrcode (QRcode payment)
	  bankCard (bankCard payment)
	  bankAccount (BankAccount payment)
	  bankTransfer (Pay with Bank Transfer)
	  bankUSSD (Pay with USSD).
	*/
	PayMethods []string `json:"payMethods,omitempty"`

	/*Payment type:
	  BalancePayment (Balance account payment)
	  BonusPayment (Bonus account payment)
	  OWealth (OWealth account payment).
	*/
	PayTypes []string `json:"payTypes,omitempty"`

	/*Product description, utf-8 encoded*/
	ProductDesc string `json:"productDesc,omitempty"`
	/*Product name, utf-8 encoded*/
	ProductName string `json:"productName,omitempty"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference,omitempty"`
	/*The address that browser go to after transaction successful.*/
	ReturnUrl string `json:"returnUrl,omitempty"`
	/*User phone number sent by merchant*/
	UserPhone string `json:"userPhone,omitempty"`
	/*The IP address requested by user, need pass-through by merchant, user Anti-phishing verification.*/
	UserRequestIp string `json:"userRequestIp,omitempty"`
}

const (
	PAYTYPES_BALANCE = "BalancePayment"
	PAYTYPES_BONUS   = "BonusPayment"
	PAYTYPES_OWEALTH = "OWealth"

	PAYMETHODS_ACCOUNT     = "account"
	PAYMETHODS_QRCODE      = "qrcode"
	PAYMETHODS_BANKCARD    = "bankCard"
	PAYMETHODS_BANKACCOUNT = "account"
)

type CashierInitializeRespDetail struct {
	/*Amount in kobo*/
	Amount string `json:"amount,omitempty"`
	/*OPay Cashier URL address, rewrite user‘s browser to it*/
	CashierUrl string `json:"cashierUrl,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Order number from OPay payment*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference,omitempty"`
	/*INITIAL PENDING SUCCESS FAIL CLOSE*/
	Status string `json:"status,omitempty"`
}

type CashierInitializeResp struct {
	Code    string                       `json:"code"`
	Data    *CashierInitializeRespDetail `json:"data"`
	Message string                       `json:"message"`
}

type CashierStatusReq struct {
	/*Order number of merchant (unique order number from merchant platform)*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of OPay payment*/
	Reference string `json:"reference,omitempty"`
}

type CashierStatusRespDetail struct {
	/*Amount in kobo*/
	Amount string `json:"amount,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Order number from OPay payment*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference,omitempty"`
	/*INITIAL PENDING SUCCESS FAIL CLOSE*/
	Status string `json:"status,omitempty"`
	/*true is support refund by original way, others not support. true false null*/
	SupportRefundOriginal *bool `json:"supportRefundOriginal"`
}

type CashierStatusResp struct {
	Code    string                   `json:"code"`
	Data    *CashierStatusRespDetail `json:"data"`
	Message string                   `json:"message"`
}

type CashierCloseReq struct {
	/*Order number of merchant (unique order number from merchant platform)*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of OPay payment*/
	Reference string `json:"reference,omitempty"`
}

type CashierCloseRespDetail struct {
	/*Order number from OPay payment*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference,omitempty"`
	/*INITIAL PENDING SUCCESS FAIL CLOSE*/
	Status string `json:"status,omitempty"`
}

type CashierCloseResp struct {
	Code    string                  `json:"code"`
	Data    *CashierCloseRespDetail `json:"data"`
	Message string                  `json:"message"`
}

type CashierRefundByOriginReq struct {
	/*Amount in kobo*/
	Amount  string `json:"amount,omitempty"`
	Country string `json:"country,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Order number of merchant original*/
	OriginalReference string `json:"originalReference,omitempty"`
	Reason            string `json:"reason,omitempty"`
	Reference         string `json:"reference,omitempty"`
	/*refundoriginal refund2bankaccount refund2opayaccount*/
	RefundType string `json:"refundType,omitempty"`
}

type CashierRefundByBankAccountReq struct {
	/*Amount in kobo*/
	Amount            string `json:"amount,omitempty"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankCode          string `json:"bankCode"`
	Country           string `json:"country,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Order number of merchant original*/
	OriginalReference string `json:"originalReference,omitempty"`
	Reason            string `json:"reason,omitempty"`
	Reference         string `json:"reference,omitempty"`
	/*refundoriginal refund2bankaccount refund2opayaccount*/
	RefundType string `json:"refundType,omitempty"`
}

type MerchantReceiver struct {
	MerchantId string `json:"merchantId"`
	Type       string `json:"type"`
}

type CashierRefundByOpayMerchantAccountReq struct {
	/*Amount in kobo*/
	Amount  string `json:"amount,omitempty"`
	Country string `json:"country,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`
	/*Order number of merchant original*/
	OriginalReference string           `json:"originalReference,omitempty"`
	Reason            string           `json:"reason,omitempty"`
	Receiver          MerchantReceiver `json:"receiver"`
	Reference         string           `json:"reference,omitempty"`
	/*refundoriginal refund2bankaccount refund2opayaccount*/
	RefundType string `json:"refundType,omitempty"`
}

type UserReceiver struct {
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
}

type CashierRefundByOpayUserAccountReq struct {
	/*Amount in kobo*/
	Amount  string `json:"amount,omitempty"`
	Country string `json:"country,omitempty"`
	/*Currency charge should be performed in. Default is NGN*/
	Currency string `json:"currency,omitempty"`

	/*Order number of merchant original*/
	OriginalReference string       `json:"originalReference,omitempty"`
	Reason            string       `json:"reason,omitempty"`
	Receiver          UserReceiver `json:"receiver"`
	Reference         string       `json:"reference,omitempty"`
	/*refundoriginal refund2bankaccount refund2opayaccount*/
	RefundType string `json:"refundType,omitempty"`
}

type CashierRefundStatusRespDetail struct {
	/*Order number of OPay payment*/
	OrderNo string `json:"orderNo,omitempty"`
	/*INITIAL PENDING SUCCESS FAIL*/
	OrderStatus string `json:"orderStatus,omitempty"`
}

type CashierRefundStatusResp struct {
	Code    string                         `json:"code"`
	Data    *CashierRefundStatusRespDetail `json:"data"`
	Message string                         `json:"message"`
}

type CashierRefundStatusReq struct {
	/*Order number of OPay payment*/
	OrderNo string `json:"orderNo,omitempty"`
	/*Order number of merchant (unique order number from merchant platform)*/
	Reference string `json:"reference,omitempty"`
}
