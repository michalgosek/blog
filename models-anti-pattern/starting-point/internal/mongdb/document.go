package mongdb

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person describes person document from person collection.
type Person struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"title,omitempty"`
	Age  int                `bson:"author,omitempty"`
}
