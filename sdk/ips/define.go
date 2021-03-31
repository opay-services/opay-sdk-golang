package ips

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

type MerchantAcquiring struct {
	Payload MerchantAcquiringPayload `json:"payload"`
	Sha512  string                   `json:"sha512"`
	Type    string                   `json:"type"`
}

type MerchantAcquiringPayload struct {
	Amount           string `json:"amount"`
	Country          string `json:"country"`
	Currency         string `json:"currency"`
	Fee              string `json:"fee"`
	Channel          string `json:"channel"`
	Refunded         bool   `json:"refunded"`
	DisplayedFailure string `json:"displayedFailure"`
	Reference        string `json:"reference"`
	UpdatedAt        string `json:"updated_at"`
	Timestamp        string `json:"timestamp"`
	InstrumentId     string `json:"instrument_id"`
	InstrumentType   string `json:"instrumentType"`
	TransactionId    string `json:"transactionId"`
	Token            string `json:"token"`
	BussinessType    string `json:"bussinessType"`
	PayChannel       string `json:"payChannel"`
	Status           string `json:"status"`
}

func (msg *MerchantAcquiring) VerfiySignature(mConf *conf.OpayMerchantConf) bool {
	payload := msg.Payload

	var flag string = "f"
	if payload.Refunded {
		flag = "t"
	}
	str := fmt.Sprintf("{Amount:\"%s\",Currency:\"%s\",Reference:\"%s\",Refunded:%s,Status:\"%s\","+
		"Timestamp:\"%s\",Token:\"%s\",TransactionID:\"%s\"}",
		payload.Amount, payload.Currency,
		payload.Reference, flag,
		payload.Status, payload.Timestamp,
		payload.Token, payload.TransactionId)

	signStr := util.SignatureSHA3512([]byte(str), mConf.GetSecretKey())

	if signStr == msg.Sha512 {
		return true
	}
	return false
}

type BettingAndAirTime struct {
	Payload BettingPayload `json:"payload"`
	Sha512  string         `json:"sha512"`
	Type    string         `json:"type"`
}

type BettingPayload struct {
	ServiceType     string `json:"serviceType"`
	FeeAmount       int    `json:"feeAmount"`
	OrderAmount     int    `json:"orderAmount"`
	OrderNo         string `json:"orderNo"`
	CreateTime      string `json:"createTime"`
	MerchantId      string `json:"merchantId"`
	OrderStatus     string `json:"orderStatus"`
	Currency        string `json:"currency"`
	PayChannel      string `json:"payChannel"`
	MerchantOrderNo string `json:"merchantOrderNo"`
	ErrorMsg        string `json:"errorMsg"`
}

func (msg *BettingAndAirTime) VerfiySignature(mConf *conf.OpayMerchantConf) bool {
	str := fmt.Sprintf("{orderNo:\"%s\",merchantOrderNo:\"%s\",merchantId:\"%s\"," +
		"orderAmount:\"%d\",serviceType:\"%s\",orderStatus:\"%s\"}", msg.Payload.OrderNo,
		msg.Payload.MerchantOrderNo, msg.Payload.MerchantId,
		msg.Payload.OrderAmount, msg.Payload.ServiceType, msg.Payload.OrderStatus)

	signStr := util.SignatureSHA3512([]byte(str), mConf.GetSecretKey())
	if signStr == msg.Sha512 {
		return true
	}

	return false
}
