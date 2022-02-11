package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	name string
	age  int
}

func StructToJSON(p interface{}) string {
	data, err := json.Marshal(p)

	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return "structToJSON ERROR !!!"
	}

	return string(data)
}

func main() {
	x := Student{
		name: "123",
		age:  12,
	}
	fmt.Println(StructToJSON(x))
}
