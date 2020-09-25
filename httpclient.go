package jpushclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	charSet                  = "UTF-8"
	contentTypeJSON          = "application/json"
	defaultConnectionTimeout = 20 //seconds
	defaultSocketTimeout     = 30 // seconds
)

// SendPostString 发送字符串
func SendPostString(url, content, authCode string) (string, error) {

	//req := Post(url).Debug(true)
	req := Post(url)
	req.SetTimeout(defaultConnectionTimeout*time.Second, defaultSocketTimeout*time.Second)
	req.Header("Connection", "Keep-Alive")
	req.Header("Charset", charSet)
	req.Header("Authorization", authCode)
	req.Header("Content-Type", contentTypeJSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.Body(content)

	return req.String()
}

// SendPostBytes 发送 bytes
func SendPostBytes(url string, content []byte, authCode string) (string, error) {

	req := Post(url)
	req.SetTimeout(defaultConnectionTimeout*time.Second, defaultSocketTimeout*time.Second)
	req.Header("Connection", "Keep-Alive")
	req.Header("Charset", charSet)
	req.Header("Authorization", authCode)
	req.Header("Content-Type", contentTypeJSON)
	req.SetProtocolVersion("HTTP/1.1")
	req.Body(content)

	return req.String()
}

// SendPostBytes2 发送 bytes
func SendPostBytes2(url string, data []byte, authCode string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Charset", charSet)
	req.Header.Add("Authorization", authCode)
	req.Header.Add("Content-Type", contentTypeJSON)
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return "", err
	}
	if resp == nil {
		return "", nil
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
