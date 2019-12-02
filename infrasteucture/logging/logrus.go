package logging

import (
	"fmt"
	"os"
	"runtime"
	"time"

	logrus "github.com/sirupsen/logrus"
	"github.com/trpg/usecases"
	"github.com/trpg/usecases/logging"
)

//LogrusLogging ローカル用
type LogrusLogging struct {
	Client *logrus.Logger
}

//NewLogrusLogging New LogrusLogging
func NewLogrusLogging() *LogrusLogging {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}

	return &LogrusLogging{
		Client: logger,
	}
}

//Error Errorレベルのアプリケーションログの出力
func (log *LogrusLogging) Error(uerr *usecases.UError) {
	//DBのレコードなしはエラーログをはかない
	if uerr.IsRecordNotFoundUError() {
		return
	}

	trace := uerr.Msg
	_, file, line, ok := runtime.Caller(1) //呼び出し元のファイル名、行数を取得
	if ok {
		trace = trace + fmt.Sprintf("[%s:%d]", file, line)
	}

	log.Client.Errorln(trace)
}

//Warning Warningレベルのアプリケーションログの出力
func (log *LogrusLogging) Warning(uerr *usecases.UError) {
	log.Client.Warningln(uerr.Msg)
}

//Info Infoレベルのアプリケーションログの出力
func (log *LogrusLogging) Info(msg string) {
	log.Client.Infoln(msg)
}

//Debug Debugレベルのアプリケーションログの出力
func (log *LogrusLogging) Debug(msg string) {
	log.Client.Debugln(msg)
}

//AccessLog ginのアクセスログの出力
func (log *LogrusLogging) AccessLog(entry *logging.AccessLogEntry) {
	log.Client.Infoln(entry)
}

//AccessSlowLog ginのアクセスログの出力(スローログ)
func (log *LogrusLogging) AccessSlowLog(entry *logging.AccessLogEntry) {
	log.Client.Warningln(entry)
}

//SQLLog gormのSQLログの出力
func (log *LogrusLogging) SQLLog(v ...interface{}) {
	msgs := v[0].([]interface{})

	sql := msgs[3].(string)

	var values string
	for i, v := range msgs[4].([]interface{}) {
		s, _ := v.(string)
		if i == 0 {
			values = s
		} else {
			values = values + "," + s
		}
	}

	executionTime := msgs[2].(time.Duration)
	file := msgs[1].(string)

	log.Client.Debugln(sql + "  [" + values + "]  " + executionTime.String() + "  " + file)
}
