package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/cars_crud_mongo/db"
	"gopkg.in/mgo.v2/bson"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

//Get all item
func GetPeople(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAllPerson()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}
	w.Write(bs)
}
func GetAllCars(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAllCars()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}
	w.Write(bs)
}

//Get one item
func GetPerson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOnePerson(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}
func GetCar(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOneCar(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}
	w.Write(bs)
}

// Delete. Removes a single item (identified by parameter) from the database.
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemovePerson(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}
	w.Write([]byte("OK," + id + " has been deleted"))
}
func DeleteCar(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemoveCar(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}
	w.Write([]byte("OK," + id + " has been deleted"))
}

//Create item. Saves an item (form data) into the database.
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Create person: "))
	var person db.People

	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Id = bson.NewObjectId()

	if err := db.CreateOnePerson(person); err != nil {
		fmt.Println("Error with insert")
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	json.NewEncoder(w).Encode(person)
}
func CreateCar(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Create Car: "))
	var car db.Cars

	_ = json.NewDecoder(req.Body).Decode(&car)
	car.Id = bson.NewObjectId()

	if err := db.CreateOneCar(car); err != nil {
		fmt.Println("Error with insert")
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	json.NewEncoder(w).Encode(car)
}

//Update item
func UpdatePerson(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Update Person: "))

	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	var person db.People
	_ = json.NewDecoder(req.Body).Decode(&person)
	rs, err := db.UpdateOnePerson(id, &person)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}
	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}
	w.Write(bs)
}
func UpdateCar(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Update Person: "))

	vars := mux.Vars(req)
	id := vars["id"]
	var car db.Cars
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	_ = json.NewDecoder(req.Body).Decode(&car)
	rs, err := db.UpdateOneCar(id, &car)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}
	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}
	//json.NewEncoder(w).Encode(&rs)
	w.Write(bs)
}


//func UpdateCar(w http.ResponseWriter, req *http.Request){
//	w.Write([]byte("Update Person: "))
//vars := mux.Vars(req)
//id := vars["id"]
//var car db.Cars
//
//err := json.NewDecoder(req.Body).Decode(&car)
//if err != nil {
//	handleError(err, "Failed to marshal data: %v", w)
//	return
//}
//
//if err = db.UpdateOneCar(id, &car); err != nil {
//	handleError(err, "Failed to read database: %v", w)
//	return
//}
//json.NewEncoder(w).Encode()
//}
