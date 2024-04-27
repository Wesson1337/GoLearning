package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	num := SumNumbers(map[int]float64{1: 1, 2: 2, 3: 3})
	fmt.Println(num)

	nod := node[int]{value: 1}
	nod2 := node[int]{value: 2}
	fmt.Println(nod)
	tre := Tree[int]{Value: 1, Left: &nod, Right: &nod2}
	fmt.Printf("%#v\n", tre)

	alig := alignmentStruct{Int: 1, Bool2: true, Bool: true, Int32: 1}
	fmt.Println(reflect.ValueOf(alig).Type().Size())

	x := 2
	a := reflect.ValueOf(x)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())
	px := d.Addr().Interface().(*int)
	fmt.Println(*px)

	x = 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)               // OK, x = 2
	rx.Set(reflect.ValueOf(3)) // OK, x = 3
	// rx.SetString("hello") // will panic: string is not int

	// rx.Set(reflect.ValueOf("hello")) // will panic: string is not int

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2) // will panic: y is an interface

	ry.Set(reflect.ValueOf(3)) // OK, y = int(3)
	// ry.SetString("hello") // will panic: string is not int

	fmt.Println(y)

	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
	fmt.Println(stdout.Type() == reflect.TypeOf(*os.Stdout))

	name := stdout.FieldByName("name")
	fmt.Println(name)

	mp := map[customStruct]int{{Name: "A", Age: 1}: 1, {Name: "B", Age: 2}: 2}
	fmt.Println(mp)

	struc := customStruct{Name: "A", Age: 1}
	strucV := reflect.ValueOf(struc)
	fmt.Println(strucV.Type().Field(0).Tag.Get("tag"))

	tes := "-3"
	val, err := strconv.ParseUint(tes, 10, 64)
	fmt.Println(val, err)

	var lol struct {
		a bool
		c int64
		b int16
		d []int
	}
	fmt.Println(unsafe.Sizeof(lol))
	fmt.Println(unsafe.Alignof(lol.c))
	fmt.Println(unsafe.Offsetof(lol.b))

	fmt.Println(uintptr(unsafe.Pointer(&lol)))
}

type customStruct struct {
	Name string `tag:"first_name"`
	Age  int
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Tree[T any] struct {
	Value       T
	Left, Right *node[T]
}

type node[T any] struct {
	value       T
	left, right *node[T]
}

type alignmentStruct struct {
	Int   int64
	Int32 int32
	Bool2 bool
	Bool  bool
}