package repository

import (
	"strconv"

	mgo "gopkg.in/mgo.v2"
)

type Repository struct {
	session *mgo.Session
}

type MasterConnection struct {
	session *mgo.Session
}

var master *MasterConnection

func (repo *Repository) Initialize(config Configuration) {
	upwd := ""
	if config.DatabaseUser != "" && config.DatabasePassword != "" {
		upwd = config.DatabaseUser + ":" + config.DatabasePassword + "@"
	}

	if config.DatabasePort == 0 {
		config.DatabasePort = 27017
	}

	url := "mongodb://" + upwd + config.DatabaseServer + ":" + strconv.Itoa(config.DatabasePort) + "/" + config.Database

	if master == nil {
		master = new(MasterConnection)
	}

	if master.session == nil {
		master.session, _ = mgo.Dial(url)
		master.session.SetMode(mgo.Monotonic, true)
	}

	repo.session = master.session.Copy()
}

func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB("").C(collection)
}

func (repo *Repository) Close() {
	repo.session.Close()
}
