package main

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	GroupName string   `json:"group_name"`
	Person    []Person `json:"persons"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData :=
		`{
		"group_name": "john_group",
		"persons": [
			{
				"name": "john1",
				"age": 100
			},
			{
				"name": "john2",
				"age": 200
			},
			{
				"name": "john3",
				"age": 300
			}
		]
	}`

	jsonBytes := ([]byte)(jsonData)
	data := new(Group)

	if error := json.Unmarshal(jsonBytes, data); error != nil {
		fmt.Println("JSON Unmarshal error: ", error)
		return
	}
	fmt.Println(data.GroupName)
	fmt.Println(data.Person[0].Name)
	fmt.Println(data.Person[1].Age)
}
