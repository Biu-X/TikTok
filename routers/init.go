package routers

import (
	"biu-x.org/TikTok/modules/web"
	v1 "biu-x.org/TikTok/routers/api/v1"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
)

func Init() {
	NewWeb()
}

func NewWeb() http.Handler {
	r := web.NewRoute()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	})
	r.Mount("/douyin", v1.Route())
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return nil
	}
	return r
}
