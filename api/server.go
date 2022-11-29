package api

import (
	v1 "github.com/SaidovZohid/note-taking/api/v1"
	"github.com/SaidovZohid/note-taking/config"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouteOptions struct {
	Cfg     *config.Config
	Storage *storage.StorageI
}

// New @title           Swagger for note api
// @version         2.0
// @description     This is a note service api.
// @host      		localhost:8080
// @BasePath  		/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouteOptions) *gin.Engine {
	router := gin.New()

	handler := v1.New(&v1.HandlerV1{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})
	apiV1 := router.Group("/v1")
	{
		apiV1.POST("/users", handler.CreateUser)
		apiV1.GET("/users/:id", handler.GetUser)
		apiV1.PUT("/users/:id", handler.UpdateUser)
		apiV1.DELETE("/users/:id", handler.DeleteUser)
		apiV1.GET("/users", handler.GetAllUsers)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
