package tu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Bruinxs/util/ut"
)

func Get_M(addr, path string, query ut.M) (ut.M, error) {
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

	resp, err := http.Get(url)
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
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
