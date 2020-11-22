package router

import (
	"movieshop/api/controller"
	"movieshop/api/service"
	"movieshop/api/util"
	"movieshop/docs"
	"movieshop/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Movie Shop APIs
// @version v1.0
// @description List of Movie APIs
// @BasePath /api

func SetupRouter() *gin.Engine {

	router := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	movieapi := router.Group("/api")
	{
		// @tag.name info
		// @tag.description Gives information of the API
		infoGroup := movieapi.Group("/info")
		{
			version := "v1.0"
			info := model.Info{version}
			infoGroup.GET("", controller.NewInfoController(info).GetInfo)
		}

		// @tag.name movie
		// @tag.description Gives movies and information
		movieGroup := movieapi.Group("/movies")
		{
			fileHelper := util.NewFileHelper()
			bookService := service.NewMovieService(fileHelper)
			bookController := controller.NewMovieController(bookService)
			movieGroup.GET("", bookController.GetAllMovies)
			movieGroup.GET(":id", bookController.GetMovieById)
		}
		movieapi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}
