package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Entity is a database object
type Entity interface {
	GetId() bson.ObjectId
	SetId(id bson.ObjectId)
}

// GetEntityList - quickly convert a slice of interface to a slice of Entity
func (repo *Repository) GetEntityList(list []interface{}) []Entity {
	entities := []Entity{}
	for _, v := range list {
		entities = append(entities, v.(Entity))
	}
	return entities
}

// GetQuery - get the query object back to do more things against it like sort, limit, etc
func (repo *Repository) GetQuery(collection string, filter bson.M) *mgo.Query {
	return repo.OpenCollection(collection).Find(filter)
}

// GetList returns a filtered list
func (repo *Repository) GetList(collection string, list interface{}, filter bson.M) error {
	q := repo.OpenCollection(collection).Find(filter)
	err := q.All(list)
	return err
}

// GetAll returns a non-filtered list
func (repo *Repository) GetAll(collection string, list interface{}) error {
	return repo.GetList(collection, list, bson.M{})
}

// Save saves a doc with id, or inserts with a new id
func (repo *Repository) Save(collection string, obj Entity) error {
	if obj.GetId() == "" {
		obj.SetId(bson.NewObjectId())
	}

	_, err := repo.OpenCollection(collection).UpsertId(obj.GetId(), obj)
	return err
}

// Remove removes a doc by id
func (repo *Repository) Remove(collection string, id bson.ObjectId) error {
	err := repo.OpenCollection(collection).RemoveId(id)
	return err
}
