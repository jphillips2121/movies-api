package dao

import (
	"github.com/jphillips2121/movies-api/models"
)

// DAO is an interface for accessing dao from json file
type DAO interface {
	GetJsonData() (*models.Movies, error)
}
