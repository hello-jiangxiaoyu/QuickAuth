package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Get 发送GET请求
func Get(url string, result any, headers ...map[string]string) error {
	var header map[string]string
	if len(headers) > 0 {
		header = headers[0]
	}
	return sendHttpRequest("GET", url, nil, result, header)
}

// Post 发送POST请求
func Post(url string, data, result any, headers ...map[string]string) error {
	var header map[string]string
	if len(headers) > 0 {
		header = headers[0]
	}
	return sendHttpRequest("POST", url, data, result, header)
}

// Put 发送PUT请求
func Put(url string, data, result any, headers ...map[string]string) error {
	var header map[string]string
	if len(headers) > 0 {
		header = headers[0]
	}
	return sendHttpRequest("PUT", url, data, result, header)
}

// Delete 发送DELETE请求
func Delete(url string, result any, headers ...map[string]string) error {
	var header map[string]string
	if len(headers) > 0 {
		header = headers[0]
	}
	return sendHttpRequest("DELETE", url, nil, result, header)
}

func sendHttpRequest(method, url string, data, result any, header map[string]string) error {
	sendBody := ""
	if data != nil {
		marshal, err := json.Marshal(data)
		if err != nil {
			return err
		}
		sendBody = string(marshal)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(sendBody))
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := certificateValidation().Do(req)
	if err != nil {
		return err
	}

	defer DeferErr(resp.Body.Close)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > http.StatusIMUsed {
		return errors.New(fmt.Sprintf("%s api err, status: %s, body: %s", method, resp.Status, string(body)))
	}

	if result != nil {
		if err = json.Unmarshal(body, result); err != nil {
			return err
		}
	}

	return nil
}

var DefaultClient = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

// 免证书校验
func certificateValidation() *http.Client {
	//InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true, 则不会校验证书以及证书中的主机名和服务器主机名是否一致。
	return DefaultClient
}
