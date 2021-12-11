package request

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Request struct {
	FromCurrency string
	ToCurrencies []string
	Amount       float64
}

func New(fromCurrency string, toCurrencies string, amount string) *Request {
	currencies := strings.Split(toCurrencies, ",")
	amountFloat, _ := strconv.ParseFloat(amount, 64)
	return &Request{fromCurrency, currencies, amountFloat}
}

func (r *Request) MarshalJSON() []byte {
	data, _ := json.Marshal(r)
	return data
}
