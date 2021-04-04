package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

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

func handleRequests(ctx context.Context) {
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
	srv := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	go func() {
		srv.ListenAndServe()
	}()
	log.Printf("server started")
	//Listen to context and do a clean termination when server shuts down.
	<-ctx.Done()
	log.Printf("server stopped")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer func() {
		cancel()
	}()
	err := srv.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Printf("server exited properly")
}


func main() {
	sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//Simple clean-up code. Creates a listener on a goroutine which notifies
	//the programm if it receives a signal from the OS. For more details, see handleRequests(ctx).
	ctx, cancel := context.WithCancel(context.Background())
    go func() {
        sig := <-sigs
        fmt.Printf("Got an %s signal. Terminating...\n", sig)
        cancel()
    }()
	handleRequests(ctx)
}
