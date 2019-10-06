package service

import (
	"bookshop/model"
	"encoding/json"
	"io/ioutil"
)

type MovieService interface {
	GetAllMovies() (model.Movies, error)
}

type movieService struct{
	fileHelper
}

func NewMovieService() MovieService {
	return movieService{}
}

func (service movieService) GetAllMovies() (model.Movies, error) {
	var movies model.Movies
	bytes, err := ioutil.ReadFile("api/repository/movies.json")
	println(bytes)
	json.Unmarshal(bytes, &movies)
	println(movies.Movies[0].Director)
	return movies, err
}

