package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[	
		{
			"first_name" : "fuad",
			"last_name" : "hasan",
			"hair_color" : "black",
			"has_dog" : true
		},
		{
			"first_name" : "mehedi",
			"last_name" : "hasan",
			"hair_color" : "red",
			"has_dog" : false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling", err)
	}

	fmt.Printf("unmarshalled %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "fuadul"
	m1.LastName = "hasan"
	m1.HairColor = "black"
	m1.HasDog = true
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "masud"
	m2.LastName = "ali"
	m2.HairColor = "red"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	//newJson, err := json.Marshal(mySlice)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
