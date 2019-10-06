package router

import (
	"bookshop/api/controller"
	"bookshop/api/service"
	"bookshop/model"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	movieapi := router.Group("api/movies")
	{
		infoGroup := movieapi.Group("/info")
		{
			version := "v1.0"
			info := model.Info{version}
			infoGroup.GET("",controller.NewInfoController(info).GetInfo)
		}
		movieGroup := movieapi.Group("/")
		{
			service := service.NewMovieService()
			booksController := controller.NewMovieController(service)
			movieGroup.GET("", booksController.GetAllMovies)
		}
	}
	return router;
}
