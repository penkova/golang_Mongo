package api

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/user/cars_crudMongo/db"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

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
	fmt.Println(id)
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

//CreateItem saves an item (form data) into the database.
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}

func CreateCar(w http.ResponseWriter, req *http.Request){
	//w.Write([]byte("Create Car"))
	//rs, err := db.CreateCar()
	//if err != nil {
	//	handleError(err, "Failed to load database items: %v", w)
	//	return
	//}
	//
	//bs, err := json.Marshal(rs)
	//if err != nil {
	//	handleError(err, "Failed to load marshal data: %v", w)
	//	return
	//}
	//w.Write(bs)
}

func UpdatePerson(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}

	rs, err := db.UpdateOnePerson(id)
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

func UpdateCar(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	rs, err := db.UpdateOneCar(id)
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

// DeleteItem removes a single item (identified by parameter) from the database.
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemovePerson(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
func DeleteCar(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemoveCar(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

