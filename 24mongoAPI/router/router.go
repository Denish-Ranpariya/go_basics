package router

import (
	"github.com/Denish-Ranpariya/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.MarkMovieAsWatched).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")
	return r
}
