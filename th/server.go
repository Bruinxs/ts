package th

import (
	"net/http/httptest"

	"strings"

	"github.com/bruinxs/mbf/route"
	"github.com/bruinxs/util/ut"
)

type ServerTester struct {
	*httptest.Server
	Path map[string]string
	Mux  *route.Mux
}

func NewServerTester() *ServerTester {
	st := &ServerTester{}
	st.Mux = route.NewMux()
	st.Mux.RegisterHook = st.register
	st.Server = httptest.NewServer(st.Mux)
	st.Path = map[string]string{}
	return st
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
	return Get(st.URL, st.Path[key], query)
}

func (st *ServerTester) Post(key string, query ut.M) (ut.M, error) {
	return Post(st.URL, st.Path[key], query)
}

func (st *ServerTester) PostJ(key string, query, body ut.M) (ut.M, error) {
	return PostJson(st.URL, st.Path[key], query, body)
}
