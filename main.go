package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/user/cars_crudMongo/api"
)

//type Item struct {
//	//ID bson.ObjectId 		`json:"id" bson:"_id"`
//	ID    string `json:"id" bson:"_id,omitempty"`
//	Value int    `json:"value"`
//}

//type Person struct{
//	_id bson.ObjectId 		`json:"id" bson:"_id"`
//	Name string 			`json:"name" bson:"name"`
//	car []Cars
//}
//type Cars struct{
//	_id bson.ObjectId 		`json:"id" bson:"_id"`
//	Model string			`json:"model" bson:"model"`
//	Age int 				`json:"age" bson:"age"`
//}


func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/api/items", api.GetAllItems).Methods("GET")
	//r.HandleFunc("/api/items/{id}", api.GetItem).Methods("GET")
	//r.HandleFunc("/api/items", api.PostItem).Methods("POST")
	//r.HandleFunc("/api/items/{id}", api.DeleteItem).Methods("DELETE")

	r.HandleFunc("/people", api.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", api.GetPerson).Methods("GET")


	r.HandleFunc("/cars", api.GetAllCars).Methods("GET")
	r.HandleFunc("/cars/{id}", api.GetCar).Methods("GET")
	//r.HandleFunc("/cars", api.PostCar).Methods("POST")
	//r.HandleFunc("/cars/{id}", api.DeleteCar).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}