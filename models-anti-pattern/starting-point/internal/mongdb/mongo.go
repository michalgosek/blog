package mongdb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// db provides underlying MongoDB implementation of order functionality.
type db struct {
	cli  mongo.Client
	conn string
}

// New returns configured MongoDB client that after successful DB connection
// can manipulate order collection content.
func New(conn string) *db {
	return &db{conn: conn}
}

// InserPerson inserts document into person collection.
func (d *db) InsertPerson(name string, age int) error {
	m := Person{
		ID:   primitive.NewObjectID(),
		Name: name,
		Age:  age,
	}
	fmt.Println(m)

	// insertion document logic..
	return nil
}
