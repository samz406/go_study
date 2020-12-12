package main

import (
	"fmt"
	"reflect"
)

func ref(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main()  {

	a := 10
	b := 20.4
	c := true
	d := "this is string"

	e := make([]int, 4, 8)

	ref(a)
	ref(b)
	ref(c)
	ref(d)
	ref(e)
}

