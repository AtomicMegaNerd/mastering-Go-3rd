package main

import (
	"fmt"
	"reflect"
)

type HappyFunObject struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	hfo := HappyFunObject{1, "F2", 3.0}
	fmt.Println("hfo:", hfo)

	r := reflect.ValueOf(&hfo).Elem()
	fmt.Println("String value:", r.String())
	hfoType := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		hfoTypeName := hfoType.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, hfoTypeName, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		}
	}

	fmt.Println("hfo:", hfo)
}
