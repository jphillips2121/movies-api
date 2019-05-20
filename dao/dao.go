package dao

import (
	"../models"
)

// DAO is an interface for accessing dao from json file
type DAO interface {
	GetJsonData() (*models.Movies, error)
}
