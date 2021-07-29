package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
)

type httpclient interface {
	DoRequest(*http.Request) ([]byte, error)
	Get(url string) ([]byte, error)
	PostForm(url string, formData map[string]string) ([]byte, error)
	PostJson(url string, body map[string]interface{}) ([]byte, error)
}

type HttpClient struct {
	*http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	if client != nil {
		return &HttpClient{client}
	}
	return &HttpClient{&http.Client{}}
}

func (h *HttpClient) DoRequest(req *http.Request) ([]byte, error) {
	resp, err := h.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client do error: %v", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (h *HttpClient) Get(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create new request error: %v", err)
	}
	return h.DoRequest(request)
}

func (h *HttpClient) PostForm(url string, formData map[string]string) ([]byte, error) {
	data := neturl.Values{}
	for k, v := range formData {
		data.Add(k, v)
	}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("create new request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return h.DoRequest(req)
}

func (h *HttpClient) PostJson(url string, body map[string]interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json Marshal error: %v", err)
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("create new request error: %v", err)
	}
	return h.DoRequest(request)
}