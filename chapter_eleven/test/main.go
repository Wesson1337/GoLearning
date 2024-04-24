package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string   `json:"name"`
	Age  []string `json:"age"`
}

func main() {
	user := User{
		Name: "Alice",
		Age:  nil,
	}
	uJson, _ := json.Marshal(user)
	fmt.Println(string(uJson))
	json.Unmarshal([]byte(uJson), &user)
	age := user.Age
	fmt.Printf("User: %#v\n", age)
}
