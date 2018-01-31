package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Item represents a sample database entity.
//type Item struct {
//	ID    string `json:"id" bson:"_id,omitempty"`
//	Value int    `json:"value"`
//}

type People struct{
	_id bson.ObjectId 		`json:"id" bson:"_id"`
	Name string 			`json:"name" bson:"name"`
	Age int 				`json:"age" bson:"age"`
	//car []Cars
}

type Cars struct{
	_id bson.ObjectId 		`json:"id" bson:"_id"`
	Model string			`json:"model" bson:"model"`
	Age int 				`json:"age" bson:"age"`
	Price int				`json:"price" bson:"price"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost:27018")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("collection")
}

func collectionPerson() *mgo.Collection {
	return db.C("peopledb")
}

func collectionCars() *mgo.Collection {
	return db.C("carsdb")
}

// GetAll returns all items from the database.
func GetAllPerson() ([]People, error) {
	res := []People{}

	if err := collectionPerson().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}
func GetAllCars()([]Cars, error){
	res:= []Cars{}

	if err := collectionCars().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

//// GetOne returns a single item from the database.
func GetOne(id string) (*People, error) {
	res := People{}

	if err := collectionPerson().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
//
//// Save inserts an item to the database.
//func Save(item People) error {
//	return collectionItems().Insert(item)
//}
//
//// Remove deletes an item from the database
//func Remove(id string) error {
//	return collectionItems().Remove(bson.M{"_id": id})
//}