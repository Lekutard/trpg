package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/trpg/docs"
	"github.com/trpg/infrasteucture/db"
	"github.com/trpg/infrasteucture/logging"
	"github.com/trpg/infrasteucture/server"
	ulogging "github.com/trpg/usecases/logging"
)

// @title 試す用 API
// @version 0.1
// @description 色々試す用のプロジェクト

// @host localhost:9800
// @BasePath /api
// @schemes http
func main() {
	gin.SetMode(gin.ReleaseMode)

	//logging
	var newLogging ulogging.Logging
	newLogging = logging.NewLogrusLogging()

	//db
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	defer db.Close()

	db.LogMode(true)
	db.SetLogger(&logging.GormLogger{Logging: newLogging})

	server.Run(db, newLogging)
}
