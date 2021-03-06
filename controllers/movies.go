package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jphillips2121/movies-api/dao"
	"github.com/jphillips2121/movies-api/models"
)

// MoviesService handles the specific functionality of accessing the movie data from the JSON.
type MoviesService struct {
	Dao dao.DAO
}

// HandleGetMovies returns all movies within the JSON.
func (service *MoviesService) HandleGetMovies(w http.ResponseWriter, req *http.Request) {

	//Returns all movies in the JSON
	movies, err := service.Dao.GetJSONData()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting movie data from json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Convert all times to UNIX times from epoch
	for index, movie := range movies.Movies {
		movies.Movies[index] = convertEpochTime(movie)
	}

	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding movies json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// HandleGetMovie returns a specific movie depending on what the id in the request is.
func (service *MoviesService) HandleGetMovie(w http.ResponseWriter, req *http.Request) {

	// Check for a movie ID in request
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		fmt.Println(fmt.Errorf("no id provided for the movie"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Converts the ID to an integer to be used later
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(fmt.Errorf("the id provided is not a number: [%v]", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returns all movies in the JSON
	movies, err := service.Dao.GetJSONData()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting movie data from json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Loops through the JSON to find valid ID and return it
	for _, movie := range movies.Movies {
		if movie.MovieId == intID {
			w.Header().Set("Content-Type", "application/json")

			err = json.NewEncoder(w).Encode(convertEpochTime(movie))
			if err != nil {
				fmt.Println(fmt.Errorf("error encoding movie json: [%v]", err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}

	// If the program reaches this then the movie ID is not present
	w.WriteHeader(http.StatusNotFound)
	fmt.Println(fmt.Errorf("the id provided is not present in the database: [%v]", id))

}

// HandleMostComments returns the user with the most comments made, and the number of comments made.
func (service *MoviesService) HandleMostComments(w http.ResponseWriter, req *http.Request) {

	// Returns all movies in the JSON
	movies, err := service.Dao.GetJSONData()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting movie data from json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a map of users with a int representing how many comment they have made
	commentCount := map[string]int{}

	// Loop through all comments on every movie
	// Add users to the map if not already present, else increment the comment count by one.
	for _, movie := range movies.Movies {
		for _, user := range movie.Comments {
			if _, ok := commentCount[user.User]; ok {
				commentCount[user.User]++
			} else {
				commentCount[user.User] = 1
			}
		}
	}

	// Loop through commentCount map and assign the commenter with the most comments to 'maxCommenter'
	maxNumber := 0
	var maxCommenter string
	for k, v := range commentCount {
		if v > maxNumber {
			maxCommenter = k
			maxNumber = v
		}
	}

	responseStruct := &models.MaxCommenterResource{
		User:     maxCommenter,
		Comments: maxNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseStruct)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding comments json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

// HandleMostLikes returns the film with the most likes on it.
func (service *MoviesService) HandleMostLikes(w http.ResponseWriter, req *http.Request) {

	// Returns all movies in the JSON
	movies, err := service.Dao.GetJSONData()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting movie data from json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mostLikedMovie := &models.MoviesResource{}

	for _, movie := range movies.Movies {
		if movie.Likes > mostLikedMovie.Likes {
			*mostLikedMovie = movie
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mostLikedMovie)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding most liked movie json: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return

}

// Converts the time stamp in the comments to a readable time.
func convertEpochTime(oldMovieResource models.MoviesResource) models.MoviesResource {
	newMovieResource := models.MoviesResource{}
	newMovieResource = oldMovieResource

	oldCommentsResource := oldMovieResource.Comments
	newCommentsResource := []models.CommentsResource{}
	for index := range oldCommentsResource {
		newCommentsResource = append(newCommentsResource, oldCommentsResource[index])
		tm, _ := strconv.ParseInt(oldCommentsResource[index].DateCreated, 10, 64)
		timenew := time.Unix(tm, 0)
		newCommentsResource[index].DateCreated = timenew.String()
	}
	newMovieResource.Comments = newCommentsResource

	return newMovieResource
}
