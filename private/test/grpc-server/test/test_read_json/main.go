package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Config struct {
	A struct {
		AA struct {
			Value string
		}
	}
	B struct {
		BB struct {
			Value string
		}
	}
}
type Animal struct {
	Name  string
	Order string
}

func Test(data []uint8) []Animal {
	//	var jsonBlob = []byte(`[
	//	{"Name": "Platypus", "Order": "Monotremata"},
	//	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	//]`)
	//
	//fmt.Println(reflect.TypeOf(jsonBlob))

	var animals []Animal
	err := json.Unmarshal(data, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	return animals
}

func main() {
	data, err := os.ReadFile("testdata.json")

	if err != nil {
		log.Fatal(err)
	}

	a := Test(data)
	fmt.Printf("%+v\n", a)
	fmt.Println(reflect.TypeOf(a))

	// var j []Config
	//
	// err := json.UnMarshal(data, &j)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// fmt.Println(j)
}
