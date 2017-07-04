package tdb

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func GetC(dbName string) func(name string) *mgo.Collection {
	up := os.Getenv("TEST_MDB")
	url := "mongodb://" + up + "@loc.mdb/" + dbName
	session, err := mgo.Dial(url)
	if err != nil {
		panic(fmt.Errorf("mgo dial to %v fail with %v", url, err))
	}
	return session.DB("").C
}
