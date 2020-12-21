package mongo

import "gopkg.in/mgo.v2"

type RoadmapController struct {
	session *mgo.Session
}

func NewRoadmapController(s *mgo.Session) *RoadmapController {
	return &RoadmapController{s}
}
