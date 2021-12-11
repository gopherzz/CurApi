package parser

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL = "https://www.xe.com/currencyconverter/convert/?Amount=%.0f&From=%s&To=%s"
)

func Parse(from string, to string, amount float64) (float64, error) {
	var ch = make(chan float64)
	go getValue(from, to, amount, ch)
	res := <-ch
	if res == -1 {
		return -1, fmt.Errorf("Error while parsing")
	}
	return res, nil
}

func getValue(from, to string, amount float64, ch chan float64) {
	httpResp, err := http.Get(fmt.Sprintf(URL, amount, from, to))
	if err != nil {
		ch <- -1
		return
	}
	defer httpResp.Body.Close()
	goqueryDoc, err := goquery.NewDocumentFromReader(httpResp.Body)
	if err != nil {
		ch <- -1
		return
	}
	value := goqueryDoc.Find("p.iGrAod").Text()
	res, err := strconv.ParseFloat(strings.Split(value, " ")[0], 64)
	if err != nil {
		ch <- -1
		return
	}
	ch <- res
}
