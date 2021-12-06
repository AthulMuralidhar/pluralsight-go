package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product2 struct {
	ID int `json:"id"`
	Manufacture string `json: "manufacture"`
	Sku string `json: "sku"`
}

func JSONTest(){

	p := Product2{
		ID: 123,
		Manufacture: "some manufacture",
		Sku: "some string",
	}

	jsonVal, err := json.Marshal(&p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonVal))
	fmt.Println()

	pJson := `{
		"id":123,
		"manufacture": "random",
		"sku": "asd2131"
	}`
	p2 := Product2{}
	err = json.Unmarshal([]byte(pJson), &p2)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Manufacture)

}