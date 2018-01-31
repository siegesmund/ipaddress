package ipaddress

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Headers struct {
	Accept                  string `json:"Accept"`
	AcceptEncoding          string `json:"Accept-Encoding"`
	AcceptLanguage          string `json:"Accept-Language"`
	Connection              string `json:"Connection"`
	Dnt                     string `json:"Dnt"`
	Host                    string `json:"Host"`
	UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
	UserAgent               string `json:"User-Agent"`
}

type Response struct {
	Args    map[string]string `json:"args"`
	Headers Headers           `json:"headers"`
	Origin  string            `json:"origin"`
}

func IpAddress() (Response, error) {

	var err error
	var response Response

	url := "http://httpbin.org/get"
	resp, err := http.Get(url)

	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return response, err
	}

	return response, err
}
