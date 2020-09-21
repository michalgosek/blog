package mongdb

import (
	"blog/models-anti-pattern/anit-pattern/internal/models"
	"blog/models-anti-pattern/anit-pattern/internal/utils"
	"fmt"

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
	m := models.Person{
		ID:   utils.UUIDGenerator(),
		Name: name,
		Age:  age,
	}

	fmt.Println(m)
	// insertion document logic and rest of the logic...
	return nil
}
