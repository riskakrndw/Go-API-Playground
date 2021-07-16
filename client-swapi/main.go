package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	Name      string `json:"name"`
	BirthYear string `json:"birth_year"`
}

func main() {
	// Dapatkan response dari swapi
	resp, err := http.Get("https://swapi.dev/api/people/1/")
	if err != nil {
		log.Fatal(err)
	}
	// Dapatkan response.Body dalam bentuk []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// tampilkan response Body dalam bentuk []byte sebagai string
	fmt.Println(string(body))
	// Decode json ke dalam interface
	var person Person
	json.Unmarshal(body, &person)
	fmt.Println("nama", person.Name)
	fmt.Println("tanggal lahir", person.BirthYear)
}