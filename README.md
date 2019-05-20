# Movies Retrieval API

An API for retrieving and displaying data about movies. 

## Requirements
In order to run this API locally you will need to install the following:

- [Go](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)

## Running the application locally

1. Clone this repository: `go get github.com/jphillips2121/movies-api`
1. From the main directory run `go run main.go`

## Endpoints

The API listens on port 8081, so all paths should be prefixed with `http://localhost:8081`

Method    | Path                                            | Description
:---------|:------------------------------------------------|:-----------
**GET**   | /movies                                         | Returns all movies on the database
**GET**  | /movies/{movie_id}                              | Returns movies details back for specific movie with relevant id
**GET**   | /comments                                       | Returns the user with the most amount of comments, and the number of these comments.
**GET**   | /likes                                          | Returns movie details back for the movie with the most amount of likes

