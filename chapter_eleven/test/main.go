package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

type User struct {
	Name string   `json:"name"`
	Age  [][]map[string]interface{} `json:"age"`
}

func main() {
	user := User{
		Name: "Alice",
		Age:  make([][]map[string]interface{}, 0),
	}
	uJson, _ := json.Marshal(user)
	stringUJson := string(uJson)
	fmt.Println(string([]byte(stringUJson)))
	json.Unmarshal([]byte(uJson), &user)
	age := user.Age
	fmt.Printf("User: %#v\n", age)

	a := testFunc(func() {
		fmt.Println("Hello")
	})
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(a).Pointer()).Name())

}

type testFunc func()