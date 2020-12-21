package main

import (
	"clydotron/go_mongo/pkg/controllers"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

type application struct {
	rmc *controllers.RoadmapController
}

func main() {

	rmc := controllers.NewRoadmapController(getSession())
	app := &application{rmc}

	fmt.Println("Listening on Port 4000")
	http.ListenAndServe(":4000", app.routes())
}

func getSession() *mgo.Session {

	//connect to our local mongo ---- @todo what about production?

	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to mongodb")
	return s
}
