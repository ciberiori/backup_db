package fx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	DeviceCode string `json:"device_code"`
	UserCode   string `json:"user_code"`
}

func CallPostQuery(urlx string, methodx string) Response {
	url := urlx
	method := methodx

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}

	data := Response{}
	json.Unmarshal(body, &data)
	return data
}
