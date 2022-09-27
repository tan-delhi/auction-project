package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Register(name string, port string, id string, time string) string {

	const myurl = "http://localhost:5000/register"

	values := map[string]string{"Name": name, "Id": id, "Port": port, "Time": time}

	json_data, err := json.Marshal(values)
	res, err := http.Post(myurl, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	message := string(body)
	return message

}
