package ips

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/util"
)

type Notify struct {
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

func (msg *Notify) VerfiySignature() bool {
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

	signStr := util.SignatureSHA3512([]byte(str))

	if signStr == msg.Sha512 {
		return true
	}
	return false
}
