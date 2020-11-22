package service

import (
	"encoding/json"
	"errors"
	"movieshop/api/util"
	"movieshop/model"

	"github.com/sirupsen/logrus"
)

type MovieService interface {
	GetAllMovies() ([]model.Movie, error)
	GetMovieById(id int) (model.Movie, error)
}

type movieService struct {
	fileHelper util.FileHelper
}

func NewMovieService(helper util.FileHelper) MovieService {
	return movieService{helper}
}

func (service movieService) GetAllMovies() ([]model.Movie, error) {
	var movieLibrary model.MovieLibrary
	bytes, err := service.fileHelper.ReadJsonFile("api/repository/movies.json")
	if err != nil {
		logrus.Errorf("failed reading the movies json file: %-v", err)
		return []model.Movie{}, err
	}
	if unmarshallErr := json.Unmarshal(bytes, &movieLibrary); unmarshallErr != nil {
		logrus.Errorf("failed unmarshalling json for movies: %-v", err)
		return []model.Movie{}, unmarshallErr
	}
	return movieLibrary.Movies, err
}

func (service movieService) GetMovieById(id int) (model.Movie, error) {
	var movieResult model.Movie
	movies, err := service.GetAllMovies()
	if err != nil {
		logrus.Errorf("failed to get movie: %-v", err)
		return model.Movie{}, err
	}
	for _, movie := range movies {
		if movie.Id == id {
			movieResult = movie
		}
	}
	if movieResult == (model.Movie{}) {
		logrus.Errorf("no movie found with id: %d", id)
		return model.Movie{}, errors.New("invalid movie id")
	}
	return movieResult, err
}
