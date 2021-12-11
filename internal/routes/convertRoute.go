package routes

import (
	"gopherzz/curapi/internal/converter"
	"gopherzz/curapi/internal/request"
	"net/http"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	if from == "" {
		from = "USD"
	}
	to := r.FormValue("to")
	if to == "" {
		to = "BYN"
	}
	amount := r.FormValue("amount")
	if amount == "" {
		amount = "1"
	}

	req := request.New(from, to, amount)
	resp := converter.Convert(req)

	w.Write(resp.ToJson())
}
