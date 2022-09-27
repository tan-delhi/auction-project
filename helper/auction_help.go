package helper

import (
	"auction/schema"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Bidder(name string, id int, port string, time int) schema.Bidder {

	var bid = schema.Bidder{Id: id, Name: name, Port: port, Time: time}
	fmt.Printf("\n")
	fmt.Printf("###################### \n")
	fmt.Printf("User registered... \n")
	fmt.Println("Name : ", name)
	fmt.Println("Id : ", id)
	fmt.Println("Port : ", port)
	fmt.Println("time : ", time, "ms")
	fmt.Printf("\n")
	return bid
}

func Bid_result(bookings []schema.Bidder) (int, string, int) {

	maxa := 0
	cnt := 0
	name := ""
	for i := 0; i < len(bookings); i++ {

		price := prices(bookings[i].Port)
		if bookings[i].Time > 2000 {
			cnt += 1
			fmt.Printf("%v Booking Time exceeded", bookings[i].Name)
			continue
		}
		fmt.Printf("Name : %v \n", bookings[i].Name)
		fmt.Printf("Price bid is %v \n \n", price)

		if price > maxa && bookings[i].Time < 2000 {
			maxa = price
			name = bookings[i].Name
		}
	}

	return maxa, name, cnt
}

func prices(port string) int {

	requestURL := fmt.Sprintf("http://localhost:%v/bid", port)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("err is %v \n", err)
		return -1
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	price, _ := strconv.Atoi(string(body))
	return price

}
