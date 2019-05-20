package controllers

import (
	"github.com/jphillips2121/movies-api/dao"
	"github.com/gorilla/mux"
)

var moviesService *MoviesService

func Register(mainRouter *mux.Router) {

	moviesService = &MoviesService{
		Dao: &dao.Json{},
	}

	mainRouter.HandleFunc("/movies", moviesService.HandleGetMovies).Methods("GET").Name("get-movies")
	mainRouter.HandleFunc("/movies/{id}", moviesService.HandleGetMovie).Methods("GET").Name("get-movie")
	mainRouter.HandleFunc("/comments", moviesService.HandleMostComments).Methods("GET").Name("get-most-comments")
	mainRouter.HandleFunc("/likes", moviesService.HandleMostLikes).Methods("GET").Name("get-most-likes")
}