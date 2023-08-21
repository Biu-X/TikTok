package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"time"

	"biu-x.org/TikTok/module/log"
	"github.com/gin-gonic/gin"
)

// LogLayout 日志layout
type LogLayout struct {
	RequestMethod string // 请求方法字段
	StatusCode    int    // 状态码
	Time          time.Time
	Metadata      map[string]interface{} // 存储自定义原数据
	Path          string                 // 访问路径
	Query         string                 // 携带query
	Body          string                 // 携带body数据
	IP            string                 // ip地址
	UserAgent     string                 // 代理
	Error         string                 // 错误
	Cost          time.Duration          // 花费时间
	Source        string                 // 来源
}

type Logger struct {
	// Filter 用户自定义过滤
	Filter func(c *gin.Context) bool
	// FilterKeyword 关键字过滤(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess 鉴权处理
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// 日志处理
	Print func(LogLayout)
	// Source 服务唯一标识
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		// 放行
		c.Next()

		cost := time.Since(start)
		layout := LogLayout{
			RequestMethod: c.Request.Method,
			StatusCode:    c.Writer.Status(),
			Time:          time.Now(),
			Path:          path,
			Query:         query,
			IP:            c.ClientIP(),
			UserAgent:     c.Request.UserAgent(),
			Error:         strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:          cost,
			Source:        l.Source,
		}

		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		if l.AuthProcess != nil {
			// 处理鉴权需要的信息
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// 自行判断key/value 脱敏等
			l.FilterKeyword(&layout)
		}
		// 自行处理日志
		l.Print(layout)
	}
}

// 高亮颜色map
var colorMap = map[string]string{
	"green":   "\033[97;42m",
	"white":   "\033[90;47m",
	"yellow":  "\033[90;43m",
	"red":     "\033[97;41m",
	"blue":    "\033[97;44m",
	"magenta": "\033[97;45m",
	"cyan":    "\033[97;46m",
	"reset":   "\033[0m",
}

// highlightString 高亮字符串
func highlightString(color string, str string) string {
	// 判断是否存在颜色，不存在返回绿色
	if _, ok := colorMap[color]; !ok {
		return colorMap["green"] + str + colorMap["reset"]
	}
	return colorMap[color] + str + colorMap["reset"]
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			v, _ := json.Marshal(layout)
			StatusMessage := layout.RequestMethod + ": " + strconv.Itoa(layout.StatusCode)
			if layout.Error == "" {
				log.Logger.Info(highlightString("green", StatusMessage) + " - " + string(v))
			} else {
				log.Logger.Error(highlightString("red", StatusMessage) + " - " + string(v))
			}
		},
		Source: "TikTok",
	}.SetLoggerMiddleware()
}
