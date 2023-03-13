package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	rec1 := Record{"Alpha", 3.1416, Secret{"rcd", "horse mane hoof grass"}}

	rec1Val := reflect.ValueOf(rec1)
	fmt.Println("String value:", rec1Val.String())

	rec1Type := rec1Val.Type()
	fmt.Printf("Type of rec1 %s has %d fields\n", rec1Type, rec1Val.NumField())

}
