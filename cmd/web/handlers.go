package main

import (
	"clydotron/go_mongo/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "helloooo")
}

func (app *application) getWorkspace(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	workspace := &models.Workspace{
		Name:    "Homework Assignment X",
		Roadmap: 101,
	}

	json.NewEncoder(w).Encode(workspace)
}

func (app *application) updateWorkspace(w http.ResponseWriter, r *http.Request) {
	//decode the inc

}
