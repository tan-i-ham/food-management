package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "Welcome home page about food management!")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message: "get method called !"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message: "post method called !"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message: "put method called !"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message: "delete method called !"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message: "not found"}`))
	}
}

// food data model
type food struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

// initial food list
type allFood []food

var foodList = allFood{
	{
		ID:    "1",
		Name:  "Banana",
		Price: "100",
	},
}

func createFood(w http.ResponseWriter, r *http.Request) {
	var newFood food
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter name and price of the food")
	}
	json.Unmarshal(reqBody, &newFood)
	foodList = append(foodList, newFood)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newFood)
}

func getAllFood(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(foodList)
}

func getFood(w http.ResponseWriter, r *http.Request) {
	foodID := mux.Vars(r)["id"]

	for _, singleFood := range foodList {
		if singleFood.ID == foodID {
			json.NewEncoder(w).Encode(singleFood)
		}
	}
}

func updateFood(w http.ResponseWriter, r *http.Request) {
	foodId := mux.Vars(r)["id"]
	var updatedFood food

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error happend in update funciton")
	}
	json.Unmarshal(reqBody, &updatedFood)

	for i, singleFood := range foodList {
		if singleFood.ID == foodId {
			singleFood.Name = updatedFood.Name
			singleFood.Price = updatedFood.Price
			foodList = append(foodList[:i], singleFood)
			json.NewEncoder(w).Encode(singleFood)
		}
	}

}

func deleteFood(w http.ResponseWriter, r *http.Request) {
	foodID := mux.Vars(r)["id"]

	for i, singleFood := range foodList {
		if singleFood.ID == foodID {
			foodList = append(foodList[:i], foodList[i+1:]...)
			fmt.Fprintf(w, "The food with ID %v has been deleted successfully", foodID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/food", createFood).Methods("POST")
	router.HandleFunc("/foods", getAllFood).Methods("GET")
	router.HandleFunc("/foods/{id}", getFood).Methods("GET")
	router.HandleFunc("/foods/{id}", updateFood).Methods("PATCH")
	router.HandleFunc("/foods/{id}", deleteFood).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
