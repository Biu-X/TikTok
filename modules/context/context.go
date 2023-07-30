package context

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type contextValuePair struct {
	key     any
	valueFn func() any
}

type Context struct {
	originCtx    context.Context
	Chi          chi.Context
	contextValue []contextValuePair

	Resp http.ResponseWriter
	Req  *http.Request
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.originCtx.Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.originCtx.Done()
}

func (ctx *Context) Err() error {
	return ctx.originCtx.Err()
}

func (ctx *Context) Value(key any) any {
	for _, pair := range ctx.contextValue {
		if pair.key == key {
			return pair.valueFn()
		}
	}
	return ctx.originCtx.Value(key)
}

func (ctx *Context) JSON(status int, content any) {

}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	c := &Context{
		originCtx: req.Context(),
		Req:       req,
		Resp:      resp,
	}
	return c
}

var Default func(next http.Handler) http.Handler

func Ctx(next http.Handler) http.Handler {
	return Default(next)
}

func Ctx1() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := NewContext(w, r)
			next.ServeHTTP(ctx.Resp, ctx.Req)
		}
		return http.HandlerFunc(fn)
	}
}

func Contexter() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := NewContext(w, req)
			next.ServeHTTP(ctx.Resp, ctx.Req)
		})
	}
}
