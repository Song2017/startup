package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	_pkg "startup/_pkg"

	"github.com/gin-gonic/gin"
)

type ResponseBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w ResponseBodyWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware(platform string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Input data
		body, _ := io.ReadAll(c.Request.Body)
		var bodyData map[string]interface{}
		// mask secret data in body
		json.Unmarshal(body, &bodyData)
		inputData := make(map[string]interface{})
		timeStart := time.Now()
		inputData["path"] = c.Request.URL.Path
		inputData["query"] = c.Request.URL.RawQuery
		inputData["body"] = bodyData
		nomadLog := _pkg.LogFormat{
			Platform:   platform,
			RemoteAddr: c.Request.RemoteAddr,
			Category:   "Request",
			Label: fmt.Sprintf(
				"%s - %s", c.Request.Method, c.Request.URL.Path),
			Input: inputData,
		}
		nomadLog.LoggerInfo()

		// Process request
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		w := &ResponseBodyWriter{
			Body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w
		c.Next()

		// Output data
		nomadLog.Latency = time.Since(timeStart).Milliseconds()
		nomadLog.Category = "Response"
		nomadLog.Input = nil
		respData := make(map[string]interface{})
		respData["status"] = w.Status()
		respData["response"] = w.Body.String()
		nomadLog.Output = respData
		nomadLog.LoggerInfo()
	}
}
