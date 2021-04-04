package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)


func checkResponseStatus(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func checkResponseCounter(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response value %d. Got %d\n", expected, actual)
    }
}

//Make sure that the API is up and running before testing.
//Checks whether the API is reachable.
func TestApiHome(t *testing.T) {
    resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		log.Fatalln(err)
	 }
	 checkResponseStatus(t, http.StatusOK, resp.StatusCode)
	 fmt.Println("checkApiHome finished")
}

//Test for increment +1 endpoint. Sends json with value=0. API should return value=1.
func TestIncrement(t *testing.T) {
	var counter Value
	counter.Value = 0
	jsonData, err := json.Marshal(counter)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer(jsonData)
    resp, err := http.Post("http://localhost:8000/increment", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusOK, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	var respCounter Value
	json.Unmarshal(respBody, &respCounter)
	checkResponseCounter(t, 1, respCounter.Value)
	fmt.Println("checkIncrement finished")
}

//Test for max-limit at increment endpoint. Sends json with value > max limit (in this case > 10000). API should return code 500 error.
func TestMaxLimit(t *testing.T) {
	var counter Value
	counter.Value = 20000
	jsonData, err := json.Marshal(counter)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer(jsonData)
    resp, err := http.Post("http://localhost:8000/increment", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusInternalServerError, resp.StatusCode)
	fmt.Println("checkMaxLimit finished")
}

//Test for decrement -1 endpoint. Sends json with value=0. API should return value=-1.
func TestDecrement(t *testing.T) {
	var counter Value
	counter.Value = 0
	jsonData, err := json.Marshal(counter)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer(jsonData)
    resp, err := http.Post("http://localhost:8000/decrement", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusOK, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	var respCounter Value
	json.Unmarshal(respBody, &respCounter)
	checkResponseCounter(t, -1, respCounter.Value)
	fmt.Println("checkDecrement finished")
}

//Test for min-limit at decrement endpoint. Sends json with value < min limit (in this case < -10000). API should return code 500 error.
func TestMinLimit(t *testing.T) {
	var counter Value
	counter.Value = -20000
	jsonData, err := json.Marshal(counter)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer(jsonData)
    resp, err := http.Post("http://localhost:8000/decrement", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusInternalServerError, resp.StatusCode)
	fmt.Println("checkMinLimit finished")
}

//Test for reset counter endpoint. API should reset counter and return counter=0.
func TestResetCounter(t *testing.T) {
    resp, err := http.Get("http://localhost:8000/reset")
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusOK, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	var respCounter Value
	json.Unmarshal(respBody, &respCounter)
	checkResponseCounter(t, 0, respCounter.Value)
	fmt.Println("checkResetValue finished")
}

//Test for return value endpoint. Sends json with value=100. API should return value=100.
func TestReturnValue(t *testing.T) {
	var counter Value
	counter.Value = 100
	jsonData, err := json.Marshal(counter)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer(jsonData)
    resp, err := http.Post("http://localhost:8000/value", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	 }
    checkResponseStatus(t, http.StatusOK, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	var respCounter Value
	json.Unmarshal(respBody, &respCounter)
	checkResponseCounter(t, 100, respCounter.Value)
	fmt.Println("checkReturnValue finished")
}