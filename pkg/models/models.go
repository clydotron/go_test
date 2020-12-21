package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Row struct {
	Tasks []Task `json:"tasks"`
}

type Lane struct {
	Name  string   `json:"name"`
	Color string   `json:"colors"`
	Rows  []Row    `json:"rows"`
	Tasks [][]Task `json:"tasks"`
}

type Roadmap struct {
	Name  string        `json:"name"`
	Lanes []Lane        `json:"lanes"`
	ID    bson.ObjectId `json:"id" bson:"_id"`
}

type Workspace struct {
	Name    string `json:"name"`
	Roadmap int    `json:"roadmap"`
}
