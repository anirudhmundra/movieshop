package controller

import (
	"bookshop/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieController interface {
	GetAllMovies(context *gin.Context)
	GetMovieById(context *gin.Context)
}

type movieController struct {
	service service.MovieService
}

func NewMovieController(service service.MovieService) MovieController {
	return movieController{service}
}

// @Tags movie
// @Description Get All Movies
// @Success 200 {object} model.Movies
// @Router / [get]
func (controller movieController) GetAllMovies(context *gin.Context){
	if movies, err := controller.service.GetAllMovies(); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, movies)
	}

}

// @Tags movie
// @Description Get Movie By Id
// @Success 200 {object} model.Movie
// @Param id path int true "Movie ID"
// @Router /{id} [get]
func (controller movieController) GetMovieById(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))
	if movies, err := controller.service.GetMovieById(id); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, movies)
	}

}
