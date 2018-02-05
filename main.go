package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/user/cars_crudMongo/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/people", api.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", api.GetPerson).Methods("GET")
	r.HandleFunc("/people", api.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", api.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", api.DeletePerson).Methods("DELETE")

	r.HandleFunc("/cars", api.GetAllCars).Methods("GET")
	r.HandleFunc("/cars/{id}", api.GetCar).Methods("GET")
	r.HandleFunc("/cars", api.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", api.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", api.DeleteCar).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}