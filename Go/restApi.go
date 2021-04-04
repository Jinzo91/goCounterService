package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Struct for request-data (json). Contains the value of our counter.
type Value struct {
    Value int `json:"value"`
}

func apiHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the API! This API provides a Counter-service (increment & decrement a value, or return current value).")
}

//Increments the counter by +1.
func incrementCounter(w http.ResponseWriter, req *http.Request) {	
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	//Parse body into Value-struct (e.g. counter). Add +1 to given counter.
	//In case no counter is given, returns 1.
	var counter Value
	json.Unmarshal(reqBody, &counter)
	//Server-sided limit of [-10000, 10000] to prevent intetger-overflow
	//and for testing purposes.
	if counter.Value > 10000 {
		log.Printf("The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, "The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		return
	}
	counter.Value = counter.Value + 1
	//Parse response as json.
	json.NewEncoder(w).Encode(counter)
}

//Decrements the counter by -1.
func decrementCounter(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	//Parse body into Value-struct (e.g. counter). Subtract -1 to given counter.
	//In case no counter is given, returns -1.
	var counter Value
	json.Unmarshal(reqBody, &counter)
	//Server-sided limit of [-10000, 10000] to prevent intetger-overflow
	//and for testing purposes.
	if counter.Value < -10000 {
		log.Printf("The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, "The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		return
	}
	counter.Value = counter.Value - 1
	//Parse response as json.
	json.NewEncoder(w).Encode(counter)
}

//Simply return the value sent by the client (because we do not have a database). 
//Reason for this solution: Rest-APIs are stateless and do not save any client data.
func returnValue(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	var counter Value
	json.Unmarshal(reqBody, &counter)
	//Server-sided limit of [-10000, 10000] to prevent intetger-overflow
	//and for testing purposes.
	if counter.Value < -10000 || counter.Value > 10000 {
		log.Printf("The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, "The number %v exceeded the allowed range of [-10000, 10000]!", counter.Value)
		return
	}
	//Parse response as json.
	json.NewEncoder(w).Encode(counter)
}

//Reset the counter to 0.
func resetValue(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var counter Value
	json.Unmarshal(reqBody, &counter)
	counter.Value = 0
	json.NewEncoder(w).Encode(counter)
}

func handleRequests() {
	//Use Gorilla/mux router for better handling of requests.
	router := mux.NewRouter()
	//Routes and their used/accepted methods.
	router.HandleFunc("/", apiHome)
	router.HandleFunc("/increment", incrementCounter).Methods("POST")
	router.HandleFunc("/decrement", decrementCounter).Methods("POST")
	router.HandleFunc("/value", returnValue).Methods("POST")
	router.HandleFunc("/reset", resetValue).Methods("GET")
	//Create handler for CORS. Needed when testing with a local web-client in browser (e.g. Chrome, Safari).
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })
	handler := c.Handler(router)
	//Listen & serve on defined port.
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func main() {
	handleRequests()
}
