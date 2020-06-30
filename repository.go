package repository

import (
	mgo "github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

type Repository struct {
	session *mgo.Session
}

type MasterConnection struct {
	session *mgo.Session
}

var master *MasterConnection

func (repo *Repository) Initialize(config Configuration) error {
	if master == nil {
		master = new(MasterConnection)
	}

	if master.session == nil {
		var err error
		master.session, err = mgo.Dial(config.ConnectionString)
		if err != nil {
			return errors.Wrapf(err, "Couldn't connect on url %s", config.ConnectionString)
		}
		master.session.SetMode(mgo.Monotonic, true)
	}

	repo.session = master.session.Copy()
	return nil
}

func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB("").C(collection)
}

func (repo *Repository) Close() {
	repo.session.Close()
}
