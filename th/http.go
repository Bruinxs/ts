package th

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Bruinxs/util"
	"github.com/Bruinxs/util/ut"
)

func request(addr, path, contentType string, query ut.M, body ut.M) (ut.M, error) {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	query_s := ""
	if len(query) > 0 {
		ss := []string{}
		for key, val := range query {
			ss = append(ss, fmt.Sprintf("%v=%v", key, val))
		}
		query_s = strings.Join(ss, "&")
	}

	var url string
	var resp *http.Response
	var err error
	switch contentType {
	case "json":
		url = addr + path + "?" + query_s
		resp, err = http.Post(url, "application/json", strings.NewReader(util.I2Json(body)))

	case "form":
		url = addr + path
		resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(query_s))

	default:
		url = addr + path + "?" + query_s
		resp, err = http.Get(url)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get(%v) response code(%v) illegal", url, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m ut.M
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		err = json.Unmarshal(data, &m)
		if err != nil {
			return nil, err
		}
	} else {
		m = ut.M{"data": string(data)}
	}

	return m, nil
}

func GP_M(addr, path string, query ut.M, body ut.M) (ut.M, error) {
	if body == nil {
		return request(addr, path, "", query, nil)
	} else {
		return request(addr, path, "json", query, body)
	}
}

func Post(addr, path string, query ut.M) (ut.M, error) {
	return request(addr, path, "form", query, nil)
}
