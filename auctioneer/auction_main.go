package main

import (
	"auction/helper"
	"auction/schema"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var bookings = []schema.Bidder{}
var current_item string

var id_map = make(map[int]int)

func formhandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	resBytes := []byte(body)
	var jsonRes map[string]interface{}

	json.Unmarshal(resBytes, &jsonRes)
	name := jsonRes["entity"].(string)
	current_item = string(name)
	id := jsonRes["id"].(string)
	fmt.Printf("Entity is %v \n", name)
	fmt.Printf("Id is %v \n \n", id)
	maxa, name, cnt := helper.Bid_result(bookings)
	if maxa == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Total registered bidders : %v \n", len(bookings))
		fmt.Fprintf(w, "Number of responses : 0")

	} else {
		fmt.Fprintf(w, "Entity %v \n", current_item)
		fmt.Fprintf(w, "Id is %v \n", id)
		fmt.Fprintf(w, " Number of successful bids : %v \n ", len(bookings)-cnt)
		fmt.Fprintf(w, "Name of bidder %v \n", name)
		fmt.Fprintf(w, "Price is %v \n", maxa)
	}

}

func getItem(w http.ResponseWriter, r *http.Request) {
	if current_item == "" {
		http.Error(w, "No item for Auction", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, current_item)
}

func register(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	resBytes := []byte(body)
	var jsonRes map[string]interface{}
	json.Unmarshal(resBytes, &jsonRes)

	id, _ := strconv.Atoi(jsonRes["Id"].(string))

	if id_map[id] == 1 {

		fmt.Fprintf(w, "Id already exists")
		return
	}
	id_map[id] = 1

	name := jsonRes["Name"].(string)
	port := jsonRes["Port"].(string)
	time, _ := strconv.Atoi(jsonRes["Time"].(string))

	var bid = helper.Bidder(name, id, port, time)
	bookings = append(bookings, bid)

	fmt.Fprintf(w, "Success")

}

func showBidders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The list of bidders are : \n")
	for i := 0; i < len(bookings); i++ {
		fmt.Fprintf(w, "Name : %v \n", bookings[i].Name)
		fmt.Fprintf(w, "Id : %v \n", bookings[i].Id)
		fmt.Fprintf(w, "Port : %v \n", bookings[i].Port)
		fmt.Fprintf(w, "Time to respond : %v ms \n ", bookings[i].Time)
		fmt.Fprintf(w, "\n")
	}

}

func main() {
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/register", register)
	http.HandleFunc("/showBidders", showBidders)
	http.HandleFunc("/get_item", getItem)

	fmt.Printf("Starting at port 5000 \n \n")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Printf("Port not started")
		log.Fatal(err)
	}

}
