package proxy

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Biu-X/TikTok/module/log"
	"github.com/gin-gonic/gin"
)

type ProxyOptions struct {
	Target      string
	PathRewrite string
}

//r.Use(proxy.HandleProxy("/tiktok", proxy.ProxyOptions{
//	Target:      "http://192.168.1.4:9000",
//	PathRewrite: "",
//}))

func HandleProxy(path string, proxyOption ProxyOptions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Index(ctx.Request.RequestURI, path) == 0 {
			client := &http.Client{}
			requestUrl := strings.Replace(ctx.Request.RequestURI, proxyOption.PathRewrite, "", -1)
			url := proxyOption.Target + requestUrl
			req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
			if err != nil {
				println(err)
				return
			}
			req.Header = ctx.Request.Header
			resp, err := client.Do(req)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			for key, value := range resp.Header {
				if len(value) == 1 {
					ctx.Writer.Header().Add(key, value[0])
				}
			}
			ctx.Status(resp.StatusCode)
			_, err = ctx.Writer.Write(body)
			if err != nil {
				log.Logger.Error(err)
				return
			}
		} else {
			ctx.Next()
		}
	}
}
