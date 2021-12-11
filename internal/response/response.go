package response

import (
	"encoding/json"
	"strconv"
)

type Response struct {
	StatusCode int        `json:"status_code"`
	Currencies []Currency `json:"currencies"`
}

func New() *Response {
	return &Response{
		Currencies: make([]Currency, 0),
	}
}

// Marshal Response to json
func (r *Response) ToJson() []byte {
	jsonBytes, _ := json.Marshal(r)
	return jsonBytes
}

type Currency struct {
	Code  string
	Value float64
}

func (c Currency) String() string {
	return c.Code + ": " + strconv.FormatFloat(c.Value, 'f', 2, 64)
}
