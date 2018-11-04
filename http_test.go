package testutil_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruinxs/com"
	testutil "github.com/bruinxs/test-util"
	"github.com/stretchr/testify/assert"
)

func TestGetAndPost(t *testing.T) {
	assert := assert.New(t)

	var out com.Map
	var key, val string

	mux := http.NewServeMux()
	mux.HandleFunc("/out", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		assert.Equal(val, r.Form.Get(key))

		data, err := json.Marshal(out)
		assert.Nil(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	ts := httptest.NewServer(mux)

	for _, f := range []func(string, string, com.Map) (com.Map, error){testutil.Get, testutil.Post} {
		for _, arg := range []struct {
			out      com.Map
			key, val string
		}{
			{com.Map{"string": "string"}, "key", "string"},
			{com.Map{"int": 1.0}, "key", "1"},
			{com.Map{"float": 3.14}, "key", "3.14"},
			{com.Map{"map": map[string]interface{}{"k": "v"}}, "key", "map"},
		} {
			out, key, val = arg.out, arg.key, arg.val
			res, err := f(ts.URL, "/out", com.Map{arg.key: arg.val})
			assert.Nil(err)
			assert.EqualValues(out, res)
		}
	}
}

func TestPostJSON(t *testing.T) {
	assert := assert.New(t)

	var out com.Map

	mux := http.NewServeMux()
	mux.HandleFunc("/out", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		bys, err := ioutil.ReadAll(r.Body)
		assert.Nil(err)
		out = nil
		err = json.Unmarshal(bys, &out)
		assert.Nil(err)
		w.Write([]byte("success"))
	}))
	ts := httptest.NewServer(mux)

	//1
	for _, arg := range []struct {
		in com.Map
	}{
		{com.Map{"string": "string"}},
		{com.Map{"int": 1.0}},
		{com.Map{"float": 3.14}},
		{com.Map{"map": map[string]interface{}{"k": "v"}}},
	} {
		res, err := testutil.PostJSON(ts.URL, "/out", com.Map{}, arg.in)
		assert.Nil(err)
		assert.Equal(res.String("data"), "success")
		assert.EqualValues(out, arg.in)
	}
}
