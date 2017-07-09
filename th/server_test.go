package th_test

import (
	"testing"

	"github.com/bruinxs/mbf/route"
	"github.com/bruinxs/ts/th"
	"github.com/bruinxs/util/ut"
)

func TestServerTester(t *testing.T) {
	st := th.NewServerTester()
	st.Mux.HandFunc("^/one/test(\\?.*)?$", func(ctx *route.SessionCtx) route.Result {
		return ctx.Success(ut.M{"msg": "ok"})
	})

	//get
	res, err := st.Get("test", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if res.StrP("msg") != "ok" {
		t.Error("ill res ", res)
		return
	}

	//post
	res, err = st.Post("test", nil)
	if err != nil {
		t.Error(err)
		return
	}
	if res.StrP("msg") != "ok" {
		t.Error("ill res ", res)
		return
	}

	//post json
	res, err = st.PostJ("test", nil, ut.M{})
	if err != nil {
		t.Error(err)
		return
	}
	if res.StrP("msg") != "ok" {
		t.Error("ill res ", res)
		return
	}
}
