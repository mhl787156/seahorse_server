package db

import (
	"fmt"

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
		panic(fmt.Sprintf("Could not connect to Mongo Database(%v). Now exiting.\n", connectURL))
	}

	newConn.Session.SetSafe(&mgo.Safe{})

	return &newConn
}

// Connect connects to rethinkDB
func Connect(dburl string) {
	if dburl == "" {
		dburl = MongoDBUrl
	}

	MongoSession = MongoConnect(dburl)
}
