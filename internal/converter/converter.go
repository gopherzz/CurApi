package converter

import (
	"fmt"
	"gopherzz/curapi/internal/parser"
	"gopherzz/curapi/internal/request"
	"gopherzz/curapi/internal/response"
	"gopherzz/curapi/internal/response/status"
	"log"
)

func Convert(req *request.Request) response.Response {
	log.Printf("%s", req.MarshalJSON())
	resp := response.Response{}
	for _, c := range req.ToCurrencies {
		value, err := parser.Parse(req.FromCurrency, c, req.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		resp.Currencies = append(resp.Currencies, response.Currency{
			Code:  req.FromCurrency + c,
			Value: value,
		})
	}
	if len(resp.Currencies) == 0 {
		resp.StatusCode = status.NotFound
		return resp
	}
	resp.StatusCode = status.OK
	return resp
}
