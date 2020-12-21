package controllers

import (
	"clydotron/go_mongo/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

var (
	dbName = "test-db"
)

type RoadmapController struct {
	session *mgo.Session
}

func NewRoadmapController(s *mgo.Session) *RoadmapController {
	return &RoadmapController{s}
}

func (rmc *RoadmapController) GetRoadmap(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	rm := models.Roadmap{}

	if err := rmc.session.DB(dbName).C("roadmaps").FindId(oid).One(&rm); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rmj, _ := json.Marshal(rm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", rmj)

	/*
		roadmap := &models.Roadmap{
			Name: "roadmap 1",
			Lanes: []models.Lane{
				models.Lane{
					Name:  "lane 1",
					Color: "orange",
					Rows: []models.Row{
						models.Row{
							Tasks: []models.Task{
								models.Task{Name: "task 1", Color: "0xff7755"},
								models.Task{"task 2", "purple"},
							},
						},
						models.Row{
							Tasks: []models.Task{
								models.Task{"task 5", "0x556699"},
								models.Task{"task 6", "purple"},
							},
						},
					},
				},
				models.Lane{
					Name:  "lane 2",
					Color: "0xcc4477",
					Rows: []models.Row{
						models.Row{
							Tasks: []models.Task{
								models.Task{"task 3", "green"},
								models.Task{"task 4", "purple"},
							},
						},
					},
				},
			},
		}
		json.NewEncoder(w).Encode(roadmap)
	*/
}

// func extractId(r *http.Request) (bson.ObjectId, bool) {
// 	id := r.URL.Query().Get(":id")
// 	if id == "" {
// 		//http.NotFound(w, r)
// 		return nil, false
// 	}

// 	if !bson.IsObjectIdHex(id) {
// 		//w.WriteHeader(http.StatusNotFound)
// 		return nil, false
// 	}

// 	oid := bson.ObjectIdHex(id)

// 	return oid, true
// }

func (rmc *RoadmapController) CreateRoadmap(w http.ResponseWriter, r *http.Request) {
	//decode the inc

	rm := models.Roadmap{}
	json.NewDecoder(r.Body).Decode(&rm)

	rm.ID = bson.NewObjectId()

	rmc.session.DB(dbName).C("roadmaps").Insert(rm)

	rmj, _ := json.Marshal(rm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", rmj)

}

func (rmc *RoadmapController) UpdateRoadmap(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	//fmt.Println(id, oid)

	rm := models.Roadmap{}
	json.NewDecoder(r.Body).Decode(&rm)
	rm.ID = oid
	fmt.Println(rm)

	err := rmc.session.DB(dbName).C("roadmaps").UpdateId(oid, &rm)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rmj, _ := json.Marshal(rm)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", rmj)
}

func (rmc *RoadmapController) DeleteRoadmap(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := rmc.session.DB(dbName).C("roadmaps").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted Roadmap: %s\n", oid)
}
