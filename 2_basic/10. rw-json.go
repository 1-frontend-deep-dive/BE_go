package goBasic

import (
	"encoding/json"
	"log"
)

const users = `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "brow",
			"has_dog": false
		},
		{
			"first_name": "John",
			"last_name": "Wick",
			"hair_color": "black",
			"has_dog": true
		}
	]
`

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func readJson() {
	unmarshal := []Person{}
	err := json.Unmarshal([]byte(users), &unmarshal)

	if err != nil {
		log.Println("Error on unmarshal json file")
	}

	log.Printf("unmarshal: %v", unmarshal)
}

func writeJson() {
	mySlice := []Person {
		{
			FirstName: "Spider",
			LastName: "Man",
			HairColor: "brown",
			HasDog: false,
		},
		{
			FirstName: "Aqua",
			LastName: "Man",
			HairColor: "black",
			HasDog: false,
		},
	}

	marshalJson, err := json.MarshalIndent(mySlice, "", "  ")

	if (err != nil) {
		log.Println("Error on marshal to json file")
	}

	log.Printf("marshal: %v", string(marshalJson))
}

func PlayWithRwJson() {
	readJson()
	writeJson()
}