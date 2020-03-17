package cbrf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type cbrfData struct {
	Date         string
	PreviousDate string
	Timestamp    string
	Valute       map[string]valute
}

type valute struct {
	Id       string `json:"ID"`
	NumCode  string
	CharCode string
	Nominal  int
	Name     string
	Value    float32
	Previous float32
}

func (val valute) String() string {
	return fmt.Sprintf("id: %s<br>"+
		"numcode: %s<br>"+
		"charcode: %s<br>"+
		"nominal%d<br>"+
		"name: %s<br>"+
		"value: %.2f<br>"+
		"previous: %.2f<br>", val.Id, val.NumCode, val.CharCode, val.Nominal, val.Name, val.Value, val.Previous)
}

func (cbrf cbrfData) String() string {
	return fmt.Sprintf(
		"Date: %v<br>"+
			"PreviousDate: %v<br>"+
			"Timestamp: %v<br>"+
			"Valute: %v<br>",
		cbrf.Date, cbrf.PreviousDate, cbrf.Timestamp, cbrf.Valute)
}

type Currency struct {
	Name     string
	Exchange float32
}

func getCurrencis() ([]byte, error) {
	response, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func GetQuotes(name string) (res string, err error) {
	data, err := getCurrencis()

	if err != nil {
		return "nil", err
	}

	cbrf := cbrfData{}

	json.Unmarshal(data, &cbrf)

	return fmt.Sprintf("%s", cbrf), nil
}
