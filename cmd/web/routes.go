package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))

	mux.Get("/api/v2/workspace/:id", http.HandlerFunc(app.getWorkspace))
	mux.Post("/api/v2/workspace/:id", http.HandlerFunc(app.updateWorkspace))

	mux.Get("/api/v2/roadmap/:id", http.HandlerFunc(app.rmc.GetRoadmap))
	mux.Post("/api/v2/roadmap", http.HandlerFunc(app.rmc.CreateRoadmap))
	mux.Post("/api/v2/roadmap/:id", http.HandlerFunc(app.rmc.UpdateRoadmap))
	//delete not exposed.
	return mux
}
