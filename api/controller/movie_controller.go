package controller

import (
	"movieshop/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
func (controller movieController) GetAllMovies(context *gin.Context) {
	movies, err := controller.service.GetAllMovies()
	if err != nil {
		logrus.Errorf("error occurred during GetAllMovies service call: %-v",
			err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	context.JSON(http.StatusOK, movies)

}

// @Tags movie
// @Description Get Movie By Id
// @Success 200 {object} model.Movie
// @Param id path int true "Movie ID"
// @Router /{id} [get]
func (controller movieController) GetMovieById(context *gin.Context) {
	id, paramErr := strconv.Atoi(context.Param("id"))
	if paramErr != nil {
		logrus.Errorf("invalid input for movie id: %-v", paramErr)
		context.AbortWithStatus(http.StatusBadRequest)
	}
	movies, err := controller.service.GetMovieById(id)
	if err != nil {
		logrus.Errorf("error occurred during GetMovieById service call: %-v",
			err)
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	context.JSON(http.StatusOK, movies)
}
