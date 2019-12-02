package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/trpg/infrasteucture/db"
	"github.com/trpg/interfaces/api/handlers"
	"github.com/trpg/interfaces/api/rolls"
	"github.com/trpg/usecases/logging"
)

//Run APIを起動する
func Run(db *db.DB, logging logging.Logging) {
	router := gin.Default()
	router.Use(handlers.GinLogger(logging), cors.Default(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api")

	//users API routing
	rollController := rolls.NewRollController(db, logging)
	getNoUsedAuth(v1, "/rolls/general", rollController.RollGeneral)

	router.Run()
}

//認証なしのGET
func getNoUsedAuth(route *gin.RouterGroup, endpoint string, controller gin.HandlerFunc) {
	route.GET(endpoint, controller)
}

/*func get(route *gin.RouterGroup, authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler, endpoint string, controller gin.HandlerFunc) {
	route.GET(endpoint, authHandler.SetFirebaseIDToContext, userHandler.SetUserToContext, controller)
}*/
