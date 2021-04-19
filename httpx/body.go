package httpx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadAndCheckSuccess(httpResp *http.Response, dest interface{}) error {
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	if !IsSuccessStatusCode(httpResp.StatusCode) {
		return &ErrorHTTPCall{Status: httpResp.StatusCode, Body: body}
	}
	return json.Unmarshal(body, dest)
}
