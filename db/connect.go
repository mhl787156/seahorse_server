package db

import (
	"gopkg.in/mgo.v2"
)

var (
	//MongoSession details
	MongoSession *DbConn
)

const (
	//MongoDBUrl Parameters
	MongoDBUrl = "localhost:27017"
)

// DBConn
type DbConn struct {
	Session       *mgo.Session
	DialInfo      *mgo.DialInfo
	ConnectionURL string
}

//MongoConnect connects to mongodb
func MongoConnect(connectURL string) *DbConn {
	var newConn DbConn
	var err error
	newConn.ConnectionURL = connectURL
	newConn.DialInfo, err = mgo.ParseURL(connectURL)
	newConn.Session, err = mgo.Dial(connectURL)

	if err != nil {
		panic("Could not connect to Mongo Database. Now exiting.\n")
	}

	newConn.Session.SetSafe(&mgo.Safe{})

	return &newConn
}

// Connect connects to rethinkDB
func Connect() {
	MongoSession = MongoConnect(MongoDBUrl)
}
