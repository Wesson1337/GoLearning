package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//!+Display

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-display
func main() {
	type Person struct {
		Name string
		Age  int
	}

	type Address struct {
		City string
		Zip  int
	}

	type Employee struct {
		Person
		Address
		Designation string
	}

	emp := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Address: Address{
			City: "New York",
			Zip:  10001,
		},
		Designation: "Software Engineer",
	}

	Display("emp", emp)

	// Output:
	// Display emp (main.Employee):
	// emp.Person = main.Person value
	// emp.Person.Name = "John Doe"
	// emp.Person.Age = 30
	// emp.Address = main.Address value
	// emp.Address.City = "New York"
	// emp.Address.Zip = 10001
	// emp.Designation = "Software Engineer"
	

	mp := map[string]int{"one": 1, "two": 2}
	Display("mp", mp)

	// Output:
	// Display mp (map[string]int):
	// mp["one"] = 1
	// mp["two"] = 2


	// Display a slice
	slice := []int{1, 2, 3}
	Display("slice", slice)

	// Output:
	// Display slice ([]int):
	// slice[0] = 1
	// slice[1] = 2
	// slice[2] = 3

}