package main

import (
	"auction/helper"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func bid(w http.ResponseWriter, r *http.Request) {

	requestURL := fmt.Sprintf("http://localhost:5000/get_item")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("err is %v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	name := string(body)

	fmt.Printf("Entity for bid is %v \n", name)

	rand.Seed(time.Now().UnixNano())
	price := strconv.Itoa(rand.Intn(3000))

	fmt.Printf("Bid price  : %v \n \n", price)

	fmt.Fprintf(w, price)

}
func main() {

	var name string
	var port string
	var id string
	var time string

	for {
		fmt.Println("Enter Name Id Port and Time")
		fmt.Scanln(&name, &id, &port, &time)

		fmt.Printf("Registering the bidder.. \n \n ")

		message := helper.Register(name, port, id, time)
		fmt.Println(message)
		if message != "Success" {
			continue
		}
		fmt.Printf("Registeration successful !! \n \n")
		fmt.Println("Name : ", name)
		fmt.Println("Id : ", id)
		fmt.Println("Port : ", port)
		fmt.Println("time : ", time, "ms")
		break
	}

	bid_port := ":" + port
	http.HandleFunc("/bid", bid)
	err := http.ListenAndServe(bid_port, nil)

	if err != nil {
		log.Fatal("Error ...")
		log.Fatal(err)
	}

}
