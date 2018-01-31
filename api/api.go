package api

import (
	"net/http"
	//"strconv"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/user/cars_crudMongo/db"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllItems returns a list of all database items to the response.
//func GetAllItems(w http.ResponseWriter, req *http.Request) {
//	rs, err := db.GetAll()
//	if err != nil {
//		handleError(err, "Failed to load database items: %v", w)
//		return
//	}
//
//	bs, err := json.Marshal(rs)
//	if err != nil {
//		handleError(err, "Failed to load marshal data: %v", w)
//		return
//	}
//	w.Write(bs)
//}

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

	rs, err := db.GetOne(id)
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

	rs, err := db.GetOne(id)
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

//
// GetItem returns a single database item matching given ID parameter.
//func GetItem(w http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	id := vars["id"]
//
//	rs, err := db.GetOne(id)
//	if err != nil {
//		handleError(err, "Failed to read database: %v", w)
//		return
//	}
//
//	bs, err := json.Marshal(rs)
//	if err != nil {
//		handleError(err, "Failed to marshal data: %v", w)
//		return
//	}
//
//	w.Write(bs)
//}
//func GetCar(w http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	id := vars["id"]
//
//	rs, err := db.GetOne(id)
//	if err != nil {
//		handleError(err, "Failed to read database: %v", w)
//		return
//	}
//
//	bs, err := json.Marshal(rs)
//	if err != nil {
//		handleError(err, "Failed to marshal data: %v", w)
//		return
//	}
//
//	w.Write(bs)
//}

// PostItem saves an item (form data) into the database.
//func PostItem(w http.ResponseWriter, req *http.Request) {
//	ID := req.FormValue("id")
//	valueStr := req.FormValue("value")
//	value, err := strconv.Atoi(valueStr)
//	if err != nil {
//		handleError(err, "Failed to parse input data: %v", w)
//		return
//	}
//
//	item := db.Item{ID: ID, Value: value}
//
//	if err = db.Save(item); err != nil {
//		handleError(err, "Failed to save data: %v", w)
//		return
//	}
//
//	w.Write([]byte("OK"))
//}

// // DeleteItem removes a single item (identified by parameter) from the database.
//func DeleteItem(w http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	id := vars["id"]
//
//	if err := db.Remove(id); err != nil {
//		handleError(err, "Failed to remove item: %v", w)
//		return
//	}
//
//	w.Write([]byte("OK"))
//}