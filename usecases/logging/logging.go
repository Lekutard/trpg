package logging

import (
	"net/http"
	"time"

	"github.com/trpg/usecases"
)

//AccessLogEntry Access Log Entry
type AccessLogEntry struct {
	Status       int
	Request      *http.Request
	ResponseSize int
	IP           string
	Latency      time.Duration
	Time         time.Time
}

//Logging Loggingのインターフェース
type Logging interface {
	Error(*usecases.UError)
	Warning(*usecases.UError)
	Info(string)
	Debug(string)
	AccessLog(*AccessLogEntry)
	AccessSlowLog(*AccessLogEntry)
	SQLLog(...interface{})
}
