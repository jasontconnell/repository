package repository

import "github.com/globalsign/mgo/bson"

type Entity interface {
	GetId() bson.ObjectId
	SetId(id bson.ObjectId)
}

// GetList returns a filtered list
func (repo *Repository) GetList(collection string, list []Entity, filter bson.M) ([]Entity, error) {
	q := repo.OpenCollection(collection).Find(filter)
	err := q.All(list)
	return list, err
}

// GetAll returns a non-filtered list
func (repo *Repository) GetAll(collection string, list []Entity) ([]Entity, error) {
	return repo.GetList(collection, list, bson.M{})
}

// Save saves a doc with id, or inserts with a new id
func (repo *Repository) Save(collection string, id bson.ObjectId, obj Entity) error {
	if id == "" {
		id = bson.NewObjectId()
	}

	_, err := repo.OpenCollection(collection).UpsertId(id, obj)
	return err
}

// Remove removes a doc by id
func (repo *Repository) Remove(collection string, id bson.ObjectId) error {
	err := repo.OpenCollection(collection).RemoveId(id)
	return err
}
