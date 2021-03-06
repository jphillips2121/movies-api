package dao

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jphillips2121/movies-api/models"
)

// JSON struct provides a service to be accessed from the dao.
type JSON struct{}

// GetJSONData returns all the JSON data from the movies.json.
func (j *JSON) GetJSONData() (*models.Movies, error) {

	// Open jsonFile and handle potential errors
	jsonFile, err := os.Open("movies.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	// Read the opened file as a byte array and handle potential errors
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	movies := &models.Movies{}

	// Unmarshal byteArray which contains jsonFile contents into 'movies' and handle potential errors
	err = json.Unmarshal(byteValue, movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
