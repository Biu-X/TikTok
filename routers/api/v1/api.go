package v1

import (
	"biu-x.org/TikTok/modules/web"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"os"
)

func Route() *web.Route {
	r := web.NewRoute()

	// 视频流接口
	r.Get("/feed/", service)

	r.Route("/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 用户注册接口
			r.Post("/register/", service)
			// 用户登录接口
			r.Post("/login/", service)
			// 用户信息信息
			r.Get("/", service)
		})
	})

	r.Route("/publish", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 投稿接口
			r.Post("/action/", service)
			// 发布列表接口
			r.Get("/list/", service)
		})
	})

	r.Route("/favorite", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 赞操作接口
			r.Post("/action/", service)
			// 喜欢列表接口
			r.Get("/list/", service)
		})
	})

	r.Route("/comment", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 评论操作接口
			r.Post("/action/", service)
			// 评论列表接口
			r.Get("/list/", service)
		})
	})

	r.Route("/relation", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 关注操作接口
			r.Post("/action/", service)
			r.Group(func(r chi.Router) {
				r.Route("/follow", func(r chi.Router) {
					// 关注列表接口
					r.Get("/list/", service)
				})
				r.Route("/follower", func(r chi.Router) {
					// 粉丝列表接口
					r.Get("/list/", service)
				})
				r.Route("/friend", func(r chi.Router) {
					// 粉丝列表接口
					r.Get("/list/", service)
				})
			})
		})
	})

	r.Route("/message", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// 发送消息接口
			r.Post("/action/", service)
			// 聊天记录接口
			r.Get("/chat/", service)
		})
	})

	r.Post("/upload/", func(w http.ResponseWriter, req *http.Request) {
		file, f, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}
		dst := "./" + f.Filename
		fmt.Println(f.Filename, f.Header)
		ff, _ := os.Create(dst)
		defer ff.Close()
		io.Copy(ff, file)
		service(w, req)
	})

	return r
}

func service(w http.ResponseWriter, req *http.Request) {
	type resp struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
		UserId     int    `json:"user_id"`
		Token      string `json:"token"`
	}

	r := &resp{
		StatusCode: 0,
		StatusMsg:  "hello",
		UserId:     1,
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJUaWt0b2siLCJleHAiOjE2OTA2MzUyODksImN1c3RvbV9jbGFpbXMiOm51bGx9.VaGM50aiQukdnxKNPYHTx0RFx583C0UUF9RYNVKfOCw",
	}
	render.Status(req, http.StatusOK)
	render.JSON(w, req, r)
}
