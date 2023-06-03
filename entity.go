package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Entity is a database object
type Entity interface {
	GetId() primitive.ObjectID
}

// GetEntityList - quickly convert a slice of interface to a slice of Entity
func (repo *Repository[T]) GetEntityList(list []interface{}) []Entity {
	entities := []Entity{}
	for _, v := range list {
		entities = append(entities, v.(Entity))
	}
	return entities
}

func (repo *Repository[T]) GetList(collection string, filter bson.M) ([]T, error) {
	client, err := repo.connect()
	if err != nil {
		return nil, fmt.Errorf("can't connect to repo. %w", err)
	}
	defer client.Disconnect(context.TODO())

	var list []T
	cursor, err := client.Database(repo.database).Collection(collection).Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("couldn't find %s %v. %w", collection, filter, err)
	}
	err = cursor.All(context.TODO(), &list)
	return list, err
}

// GetAll returns a non-filtered list
func (repo *Repository[T]) GetAll(collection string) ([]T, error) {
	return repo.GetList(collection, bson.M{})
}

// Save saves a doc with id, or inserts with a new id
func (repo *Repository[T]) Save(collection string, obj T) (primitive.ObjectID, error) {
	id := obj.GetId()
	if id == primitive.NilObjectID {
		id = primitive.NewObjectID()
	}

	client, err := repo.connect()
	if err != nil {
		return id, fmt.Errorf("can't connect to repo. %w", err)
	}
	defer client.Disconnect(context.TODO())

	opts := options.Replace().SetUpsert(true)
	_, err = client.Database(repo.database).Collection(collection).ReplaceOne(context.TODO(), bson.M{"_id": obj.GetId()}, obj, opts)
	return id, err
}

// Remove removes a doc by id
func (repo *Repository[T]) Remove(collection string, id primitive.ObjectID) error {
	client, err := repo.connect()
	if err != nil {
		return fmt.Errorf("can't connect to repo. %w", err)
	}
	defer client.Disconnect(context.TODO())

	opts := options.Delete()
	_, err = client.Database(repo.database).Collection(collection).DeleteOne(context.TODO(), id, opts)
	return err
}
