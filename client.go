package deribit

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	EndpointTest = "https://test.deribit.com"
	EndpointProd = "https://www.deribit.com"
)

type Client struct {
	URL      string
	Client   *http.Client
	ClientId string
	Secret   string
}

func New(endpoint, clientId, secret string) *Client {
	return &Client{
		URL:      endpoint,
		Client:   &http.Client{},
		ClientId: clientId,
		Secret:   secret,
	}
}

func NewWithProxy(endpoint, clientId, secret, httpProxy string) *Client {
	dbClient := New(endpoint, clientId, secret)
	if httpProxy != "" {
		proxyAddress, _ := url.Parse(httpProxy)
		dbClient.Client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyAddress),
		}
	}
	return dbClient
}

func (c *Client) Get(path string, query map[string]interface{}) ([]byte, error) {
	url := c.URL + path + "?" + toQueryString(query)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if strings.Contains(path, "private") {
		s := c.ClientId + ":" + c.Secret
		authorization := "Basic " + base64.URLEncoding.EncodeToString([]byte(s))
		request.Header.Add("Authorization", authorization)
	}
	request.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func valueOf(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}
func getPointer(v interface{}) interface{} {
	vv := valueOf(v)
	if vv.Kind() == reflect.Ptr {
		return v
	}
	return reflect.New(vv.Type()).Interface()
}

func (c *Client) GetAndUnmarshalJson(path string, query map[string]interface{}, v interface{}) error {
	jsonData, err := c.Get(path, query)
	if err != nil {
		return err
	}
	pointer := getPointer(v)
	return json.Unmarshal(jsonData, &pointer)
}

func toQueryString(queryMap map[string]interface{}) (qs string) {
	if queryMap == nil {
		return ""
	}
	for k, v := range queryMap {
		var value string
		i, ok := v.(int)
		if ok {
			value = strconv.Itoa(i)
		}
		switch v.(type) {
		case int:
			value = strconv.Itoa(i)
			break
		case string:
			value = fmt.Sprintf("%v", v)
		}
		if len(value) == 0 {
			continue
		}
		qs += k + "=" + value + "&"
	}
	qs = qs[:len(qs)-1]
	return
}
