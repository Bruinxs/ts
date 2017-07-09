package th

import (
	"net/http/httptest"

	"strings"

	"github.com/bruinxs/mbf/route"
	"github.com/bruinxs/util"
	"github.com/bruinxs/util/ut"
)

type ServerTester struct {
	*httptest.Server
	Path        map[string]string
	Mux         *route.Mux
	RespChecker func(resp ut.M) (ut.M, error)
}

func NewServerTester() *ServerTester {
	st := &ServerTester{}
	st.Mux = route.NewMux()
	st.Mux.RegisterHook = st.register
	st.Server = httptest.NewServer(st.Mux)
	st.Path = map[string]string{}
	st.RespChecker = defaultChecker
	return st
}

func defaultChecker(resp ut.M) (ut.M, error) {
	if resp == nil {
		return nil, util.Err("arg resp is nil")
	}
	if !resp.Exist("code") || resp.Int("code") != 0 {
		return resp, util.Err("resp illegal data -> %v", util.I2Json(resp))
	}
	return resp.Map("data"), nil
}

func (st *ServerTester) register(_, pattern string, _ route.Handle) {
	pattern = strings.Trim(pattern, "^(\\?).*$")
	if pattern == "" {
		return
	}
	ps := strings.Split(pattern, "/")
	key := ps[len(ps)-1]
	st.Path[key] = pattern
}

func (st *ServerTester) Get(key string, query ut.M) (ut.M, error) {
	resp, err := Get(st.URL, st.Path[key], query)
	if err != nil {
		return nil, err
	}
	if st.RespChecker != nil {
		return st.RespChecker(resp)
	} else {
		return resp, nil
	}
}

func (st *ServerTester) Post(key string, query ut.M) (ut.M, error) {
	resp, err := Post(st.URL, st.Path[key], query)
	if err != nil {
		return nil, err
	}
	if st.RespChecker != nil {
		return st.RespChecker(resp)
	} else {
		return resp, nil
	}
}

func (st *ServerTester) PostJ(key string, query, body ut.M) (ut.M, error) {
	resp, err := PostJson(st.URL, st.Path[key], query, body)
	if err != nil {
		return nil, err
	}
	if st.RespChecker != nil {
		return st.RespChecker(resp)
	} else {
		return resp, nil
	}
}
