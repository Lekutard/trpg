package logging

import "github.com/trpg/usecases/logging"

// GormLogger Gormのloggerをラップする
type GormLogger struct {
	Logging logging.Logging
}

// Print Gormのログをloggerに送る
func (gl *GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		gl.Logging.SQLLog(v)
	}
}
