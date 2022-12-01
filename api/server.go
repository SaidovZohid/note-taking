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
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
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
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})
	apiV1 := router.Group("/v1")
	router.Static("/medias", "./media")
	{
		apiV1.POST("/users", handler.AuthMiddileWare, handler.CreateUser)
		apiV1.GET("/users/:id", handler.AuthMiddileWare, handler.GetUser)
		apiV1.PUT("/users/:id", handler.AuthMiddileWare, handler.UpdateUser)
		apiV1.DELETE("/users/:id", handler.AuthMiddileWare, handler.DeleteUser)
		apiV1.GET("/users", handler.GetAllUsers)

		apiV1.POST("/notes", handler.AuthMiddileWare, handler.CreateNote)
		apiV1.GET("/notes/:id", handler.AuthMiddileWare, handler.GetNote)
		apiV1.PUT("/notes/:id", handler.AuthMiddileWare, handler.UpdateNote)
		apiV1.DELETE("/notes/:id", handler.AuthMiddileWare, handler.DeleteNote)
		apiV1.GET("/notes", handler.GetAllNotes)

		apiV1.POST("/auth/register", handler.Register)
		apiV1.POST("/auth/verify", handler.Verify)
		apiV1.POST("/auth/login", handler.Login)

		apiV1.POST("/file-upload", handler.AuthMiddileWare, handler.UploadFile)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
