package middleware

import (
	"blog-go-gin/helper"
	"blog-go-gin/logging"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"time"
)

//日志中间件
type loggerFunc func(format string, args ...interface{})

//获取调用方
func GetPeerNameByPath(path string) string {
	switch path {
	case "/":
		return ""
	default:
		return "client"
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求相关信息
		path := c.Request.URL.Path
		//过滤健康检查请求
		//if path == "/" {
		//	return
		//}
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		header := c.Request.Header
		reqData := ""
		if method == "POST" {
			data, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			_ = helper.HideSensitiveInfo(&data, c.Request.URL.Path) // 隐藏敏感信息
			maxLen := int(math.Min(float64(len(data)), float64(2048)))
			reqData = string(data)[:maxLen]
		} else if method == "GET" {
			params := c.Request.URL.Query()
			data, err := json.Marshal(params)
			if err != nil {
				reqData = params.Encode()
			} else {
				reqData = string(data)
			}
		}
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		//请求相关信息
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		format := "|%3d|%13v|%15s|%s|%s|%s|%s|%s|"
		var logger loggerFunc
		if errorMessage == "" {
			logger = logging.Entry.WithFields(logrus.Fields{
				"cost_time":  latency,
				"peer_name":  GetPeerNameByPath(c.Request.URL.Path),
				"req_method": method,
				"req_uri":    c.Request.URL.Path,
				"real_ip":    clientIP,
				"http_code":  statusCode,
			}).Infof
		} else {
			logger = logging.Entry.WithFields(logrus.Fields{
				"error_code": statusCode,
				"error_msg":  errorMessage,
			}).Errorf
		}
		logger(format,
			statusCode,
			latency,
			clientIP,
			method,
			path,
			header,
			reqData,
			errorMessage,
		)
	}
}
