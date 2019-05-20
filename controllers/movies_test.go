package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/jphillips2121/movies-api/models"

	"github.com/golang/mock/gomock"
	"github.com/jphillips2121/movies-api/dao"
	. "github.com/smartystreets/goconvey/convey"
)

var defaultComment1 = models.CommentsResource{
	User:        "test1",
	Message:     "message1",
	DateCreated: "1294012800",
}

var defaultComment2 = models.CommentsResource{
	User:        "test2",
	Message:     "message2",
	DateCreated: "1294012800",
}

var defaultComment3 = models.CommentsResource{
	User:        "test1",
	Message:     "message3",
	DateCreated: "1294012800",
}

var defaultMovie = models.MoviesResource{
	MovieId:       1,
	Title:         "DefaultMovie",
	Description:   "DefaultDesc",
	Producer:      "DefaultProducer",
	AvailableIn3D: true,
	AgeRating:     "AgeRating",
	Likes:         10,
	Comments:      []models.CommentsResource{defaultComment1, defaultComment2, defaultComment3},
}

var defaultAllMoviesResource = models.Movies{
	Movies: []models.MoviesResource{defaultMovie},
}

func createMockMoviesService(mockDao *dao.MockDAO) MoviesService {
	return MoviesService{
		Dao: mockDao,
	}
}

func TestMoviesService_HandleGetMovies(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error Getting Movie Data from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(nil, fmt.Errorf("error getting JSON"))

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovies(w, req)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("Successfully Returns Movie Data from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(&defaultAllMoviesResource, nil)

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovies(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
	})
}

func TestMoviesService_HandleGetMovie(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("No ID Provided For Movie", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovie(w, req)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("ID cannot be converted to an Integer", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)

		req := httptest.NewRequest("Get", "/test", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovie(w, req)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("Error Getting Movie Data from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(nil, fmt.Errorf("error getting JSON"))

		req := httptest.NewRequest("Get", "/test", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovie(w, req)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("Movie ID is not present in the Database", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(&defaultAllMoviesResource, nil)

		req := httptest.NewRequest("Get", "/test", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "2"})
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovie(w, req)

		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Successfully return movie from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(&defaultAllMoviesResource, nil)

		req := httptest.NewRequest("Get", "/test", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()
		mockMoviesService.HandleGetMovie(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
	})
}

func TestMoviesService_HandleMostComments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error Getting Movie Data from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(nil, fmt.Errorf("error getting JSON"))

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleMostComments(w, req)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("Successfully returns the User with the most comments", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(&defaultAllMoviesResource, nil)

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleMostComments(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
	})

}

func TestMoviesService_HandleMostLikes(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error Getting Movie Data from JSON", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(nil, fmt.Errorf("error getting JSON"))

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleMostLikes(w, req)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("Successfully returns the User with the most comments", t, func() {
		mock := dao.NewMockDAO(mockCtrl)
		mockMoviesService := createMockMoviesService(mock)
		mock.EXPECT().GetJSONData().Return(&defaultAllMoviesResource, nil)

		req := httptest.NewRequest("Get", "/test", nil)
		w := httptest.NewRecorder()
		mockMoviesService.HandleMostLikes(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
	})
}
