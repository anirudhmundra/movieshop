package service

import (
	"bookshop/api/util"
	"bookshop/model"
	"encoding/json"
)

type MovieService interface {
	GetAllMovies() (model.Movies, error)
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
		println(err.Error())
		return model.Movies{}, err
	}
	json.Unmarshal(bytes, &movies)
	return movies, err
}

