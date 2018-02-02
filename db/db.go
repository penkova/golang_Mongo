package db

import (
	"log"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type People struct{
	Id bson.ObjectId 		`json:"id" bson:"_id"`
	Name string 			`json:"name" bson:"name"`
	Age int 				`json:"age" bson:"age"`
	//car []Cars
	// id = car_id
}

type Cars struct{
	//car_id bson.ObjectId 	`json:"id" bson:"_id"`
	Id bson.ObjectId 		`json:"id" bson:"_id,omitempty"`
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
	db = session.DB("crudmongo" )
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
func GetOneCar(id string) (*Cars, error) {
	res := Cars{}

	if err := collectionCars().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func GetOnePerson(id string) (*People, error) {
	res := People{}
	//err := peopleColl.FindId(bson.ObjectIdHex(id)).One(&person)
	fmt.Println(id)

	if err := collectionPerson().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an item to the database.
func CreateCar(item Cars) error {
	return collectionCars().Insert(item)
}

// Remove deletes an item from the database
func RemovePerson(id string) error {
	return collectionPerson().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

func RemoveCar(id string) error {
	return collectionCars().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}