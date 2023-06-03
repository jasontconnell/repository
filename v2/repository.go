package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[T Entity] struct {
	connstr  string
	database string
}

func (repo *Repository[T]) Initialize(config Configuration) {
	repo.connstr = config.ConnectionString
	repo.database = config.Database
}

func (repo *Repository[T]) connect() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(repo.connstr))
}
