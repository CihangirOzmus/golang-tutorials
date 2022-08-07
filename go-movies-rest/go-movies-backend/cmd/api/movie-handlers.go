package main

import (
	"errors"
	"github.com/cozmus/golang-tutorials/go-movies-rest/go-movies-backend/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
	}

	app.logger.Println("id is", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Some desc",
		Year:        2021,
		ReleaseDate: time.Date(2022, 02, 02, 0, 0, 0, 0, time.Local),
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJson(w, http.StatusOK, movie, "movie")

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
