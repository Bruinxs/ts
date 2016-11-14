package tu

import (
	"encoding/json"
	"github.com/Bruinxs/util"
	"github.com/Bruinxs/util/ut"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet_M(t *testing.T) {
	http.HandleFunc("/p1/p2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := ut.M{"s1": "str", "i1": 1, "f1": 3.14, "o1": ut.M{"key": "val"}, "a1": []string{"fake"}}
		err := r.ParseForm()
		if err != nil {
			t.Error(err)
			return
		}
		for k, v := range r.Form {
			m[k] = v[0]
		}

		data, err := json.Marshal(m)
		if err != nil {
			t.Error(err)
			return
		}
		w.Write(data)
	}))

	ts := httptest.NewServer(http.DefaultServeMux)
	res, err := Get_M(ts.URL, "p1/p2", ut.M{"s2": "string", "i2": 10, "f2": 5.32})
	if err != nil {
		t.Error(err)
		return
	}

	if g, w := res.StrV("s1"), "str"; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}
	if g, w := res.StrV("s2"), "string"; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	if g, w := res.IntV("i1"), 1; g != w {
		t.Errorf("got(%v) != %v, res(%v)", g, w, util.I2Json(res))
		return
	}
	if g, w := res.IntV("i2"), 10; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	if g, w := res.FloatV("f1"), 3.14; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}
	if g, w := res.FloatV("f2"), 5.32; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}
}
