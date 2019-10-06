package controller

import (
	"bookshop/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieController interface {
	GetAllMovies(c *gin.Context)
}

type movieController struct {
	service service.MovieService
}

func NewMovieController(service service.MovieService) MovieController {
	return movieController{service}
}
func (controller movieController) GetAllMovies(context *gin.Context){
	if movies, err := service.NewMovieService().GetAllMovies(); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		println("Inside movies")
		context.JSON(http.StatusOK, movies)
	}

}
