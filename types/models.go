package types
import "go.mongodb.org/mongo-driver/bson/primitive"
type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string `bson:"first_name", json:first_name"`
	Lastname string `bson:"last_name", json:last_name"`
	Position string `bson:"position", json:position"`
	Salary float32  `bson:"salary", json:salary"`
}