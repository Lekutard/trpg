package handlers

import (
	"bytes"
	"fmt"
	"time"

	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/trpg/usecases"
	"github.com/trpg/usecases/logging"
	"github.com/trpg/usecases/utils"
)

// GinLogger Gin Logger
func GinLogger(l logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := utils.NowTime()

		requestBody := storeRequestBody(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		c.Next()

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		end := utils.NowTime()
		latency := end.Sub(start)

		defer finalizePanic(c, l, latency, end)

		entry := &logging.AccessLogEntry{
			Status:       c.Writer.Status(),
			Request:      c.Request,
			ResponseSize: c.Writer.Size(),
			IP:           c.ClientIP(),
			Latency:      latency,
			Time:         end,
		}

		if latency.Seconds() > 3 {
			l.AccessSlowLog(entry)
		} else {
			l.AccessLog(entry)
		}
	}
}

// storeRequestBody store RequestBody
func storeRequestBody(rc io.ReadCloser) (b []byte) {
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		panic(err)
	}
	return b
}

func finalizePanic(c *gin.Context, l logging.Logging, latency time.Duration, end time.Time) {
	if r := recover(); r != nil {
		errMsg := fmt.Sprintf("%v", r)
		uerr := usecases.GetUError(usecases.Internalexception, errMsg)
		l.Warning(uerr)

		entry := &logging.AccessLogEntry{
			Status:       c.Writer.Status(),
			Request:      c.Request,
			ResponseSize: c.Writer.Size(),
			IP:           c.ClientIP(),
			Latency:      latency,
			Time:         end,
		}
		l.AccessLog(entry)
	}
}
