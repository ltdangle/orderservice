package clients

import (
	"net/url"
	"orders/helpers"
)

const (
	GetProductByIdUrl    = "http://localhost:8081"
	GetProductByIdMethod = "GET"
)

func GetProductById(id string) (error, string) {
	url, err := url.Parse(GetProductByIdUrl)
	if err != nil {
		return err, ""
	}

	q := url.Query()
	q.Add("id", id)
	url.RawQuery = q.Encode()

	_, err = helpers.MakeHttpRequest(url.String(), GetProductByIdMethod)
	if err != nil {
		return err, ""
	}

	jsonString := `{
	  "Title": "one more item",
	  "Description": "description",
	  "Price": 20
	}`

	return nil, jsonString
}
