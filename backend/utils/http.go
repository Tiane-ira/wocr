package utils

import (
	"io"
	"net/http"
	"strings"
)

func PostWithForm(url string, params, body map[string]string) ([]byte, error) {
	headers := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}
	url = url + concatParam(params)
	data, err := doHttp(url, "POST", headers, concatBody(body))
	if err != nil {
		return data, err
	}
	return data, nil
}

func concatParam(params map[string]string) string {
	res := ""
	if len(params) > 0 {
		res = "?" + concatBody(params)
	}
	return res
}

func concatBody(params map[string]string) string {
	res := ""
	if len(params) > 0 {
		index := 0
		for k, v := range params {
			if index == 0 {
				res = k + "=" + v
			} else {
				res = res + "&" + k + "=" + v
			}
			index++
		}
	}
	return res
}

func Get(url string, headers map[string][]string, params map[string]string) (data []byte, err error) {
	url = url + concatParam(params)
	return doHttp(url, "POST", headers, "")
}

func Post(url string, headers map[string][]string, body string) (data []byte, err error) {
	return doHttp(url, "POST", headers, body)
}

func doHttp(url string, method string, headers map[string][]string, body string) (data []byte, err error) {
	var bodyR io.Reader
	if body != "" {
		bodyR = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, url, bodyR)
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header[k] = v
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
