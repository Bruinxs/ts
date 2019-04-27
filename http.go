package testutil

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/bruinxs/com"
)

var DefaultClient = &Client{http.DefaultClient}

type Client struct {
	*http.Client
}

func NewClient() *Client {
	cookieJar, _ := cookiejar.New(nil)
	return &Client{&http.Client{Jar: cookieJar}}
}

func (c *Client) do(host, path, method, contentType string, query com.Map, body io.Reader) (com.Map, error) {
	if c == nil {
		c = DefaultClient
	}

	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}

	url := host + path
	if len(query) > 0 {
		kvs := []string{}
		for k, v := range query {
			kvs = append(kvs, fmt.Sprintf("%s=%v", k, v))
		}
		url += fmt.Sprintf("?%s", strings.Join(kvs, "&"))
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request url %s response code %d is illegal, body data: %s", url, resp.StatusCode, bys)
	}

	result := make(com.Map)
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		err = json.Unmarshal(bys, &result)
		if err != nil {
			return nil, err
		}
	} else {
		result.Set("data", bys)
	}
	return result, nil
}

func (c *Client) Get(host, path string, query com.Map) (com.Map, error) {
	return c.do(host, path, "GET", "", query, nil)
}

func (c *Client) Post(host, path string, query com.Map, forms ...com.Map) (com.Map, error) {
	var body io.Reader
	if len(forms) > 0 {
		kvs := []string{}
		for k, v := range forms[0] {
			kvs = append(kvs, fmt.Sprintf("%s=%v", k, v))
		}
		body = strings.NewReader(strings.Join(kvs, "&"))
	}
	return c.do(host, path, "POST", "application/x-www-form-urlencoded", query, body)
}

func (c *Client) PostJSON(host, path string, query com.Map, body interface{}) (com.Map, error) {
	return c.do(host, path, "POST", "application/json", query, strings.NewReader(com.Json(body)))
}

func Get(host, path string, query com.Map) (com.Map, error) {
	return DefaultClient.Get(host, path, query)
}

func Post(host, path string, query com.Map) (com.Map, error) {
	return DefaultClient.Post(host, path, query)
}

func PostJSON(host, path string, query com.Map, body interface{}) (com.Map, error) {
	return DefaultClient.PostJSON(host, path, query, body)
}
