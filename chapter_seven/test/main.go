package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

const debug = false

func main() {
	var buf io.Writer
	fmt.Printf("%v\n", buf)
	if debug {
		buf = new(bytes.Buffer) // Накопления вывода
	}
	fmt.Printf("%v\n", buf)
	f(buf) // Внимание: тонкая ошибка!
	if debug {
		// Вывод buf
		fmt.Println(buf)
	}

	person := &Person{
		name: "Alice",
		age:  23,
	}
	fmt.Println(person.name)
	changePerson(person)
	fmt.Println(person.name)

	slice := StringSlice{"Bob", "Alice", "David"}
	sort.Sort(slice)
	fmt.Println(slice)

	f := funcT(func(test *string) {
		fmt.Println("here func")
		fmt.Println(*test)
	})
	str := "Hello"
	f.hello(&str)

	mp := make(map[string]string)
	mp["name"] = "Alice"
	mp["age"] = "23"
	mp["name"] = "Bob"
	fmt.Println(mp["name"])

	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("%v\n", rw)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", rw)
	rw = w.(io.ReadWriter)

	type Test struct {
		Hello string `json:"hello"`
	}

	hey := `{"hello": 1}`

	test := Test{}
	err := json.Unmarshal([]byte(hey), &test)
	fmt.Println(err)
	fmt.Printf("%v\n", test)

}

func f(out io.Writer) {
	// ... некоторые действия ...
	fmt.Printf("value - %v\n", out)
	fmt.Println(out)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

type Person struct {
	name string
	age  uint8
}

func changePerson(person *Person) {
	*person = Person{
		name: "David",
		age:  30,
	}

	person = &Person{
		name: "Bob",
		age:  25,
	}
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type funcT func(test *string)

func (f funcT) hello(test *string) {
	fmt.Println("here hello")
	f(test)
}
