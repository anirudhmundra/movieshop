package service

import (
	"bookshop/api/util"
	"bookshop/model"
	"encoding/json"
)

type MovieService interface {
	GetAllMovies() (model.Movies, error)
	GetMovieById(id int) (model.Movie, error)
}

type movieService struct{
	fileHelper util.FileHelper
}

func NewMovieService(helper util.FileHelper) MovieService {
	return movieService{helper}
}

func (service movieService) GetAllMovies() (model.Movies, error) {
	var movies model.Movies
	bytes, err := service.fileHelper.ReadJsonFile("api/repository/movies.json")
	if err != nil {
		return model.Movies{}, err
	}
	json.Unmarshal(bytes, &movies)
	return movies, err
}

func (service movieService) GetMovieById(id int) (model.Movie, error) {
	var movies model.Movies
	var movieResult model.Movie
	bytes, err := service.fileHelper.ReadJsonFile("api/repository/movies.json")
	json.Unmarshal(bytes, &movies)
	for _, movie := range movies.Movies{
		if movie.Id == id{
			movieResult = movie
		}
	}
	if err != nil {
		return model.Movie{}, err
	}
	return movieResult, err

}

