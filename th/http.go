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

func GP_M(addr, path string, query ut.M, body ut.M) (ut.M, error) {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	query_s := ""
	if len(query) > 0 {
		ss := []string{}
		for key, val := range query {
			ss = append(ss, fmt.Sprintf("%v=%v", key, val))
		}
		query_s = "?" + strings.Join(ss, "&")
	}
	url := addr + path + query_s

	var resp *http.Response
	var err error
	if body == nil {
		resp, err = http.Get(url)
	} else {
		resp, err = http.Post(url, "application/json", strings.NewReader(util.I2Json(body)))
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
