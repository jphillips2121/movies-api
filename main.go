package main

import (
	"github.com/jphillips2121/movies-api/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create router
	mainRouter := mux.NewRouter()
	controllers.Register(mainRouter)

	fmt.Println("Starting movies-api")

	err := http.ListenAndServe(":8081", mainRouter)
	if err != nil {
		fmt.Println(err)
	}

}

