package models

// Movies is an array of all the movies that can be returned
type Movies struct {
	Movies []MoviesResource `json:"movies"`
}

// MoviesResource is the JSON data about a movie that will be returned in the response.
type MoviesResource struct {
	MovieId       int                `json:"movie_id"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Producer      string             `json:"producer"`
	AvailableIn3D bool               `json:"available_in_3d"`
	AgeRating     string             `json:"age_rating"`
	Likes         int                `json:"likes"`
	Comments      []CommentsResource `json:"comments"`
}

// CommentsResource is the JSON data for a comment on a movie.
type CommentsResource struct {
	User        string `json:"user"`
	Message     string `json:"message"`
	DateCreated string `json:"dateCreated"`
	Like        int    `json:"like"`
}

// MaxCommenterResource is the JSON data returned for the commenter with the most comments.
type MaxCommenterResource struct {
	User     string `json:"user"`
	Comments int    `json:"comments"`
}
