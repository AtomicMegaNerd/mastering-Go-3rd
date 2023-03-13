package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Breed string
}

func getIntThing() interface{} {
	return 20
}

func getFloatThing() interface{} {
	return 4.2
}

func getDogThing() interface{} {
	return Dog{"Makwa", "Schnoodle"}
}

func getPersonThing() interface{} {
	return Person{"Chris", 47}
}

func doThingWithThing(thing interface{}) {
	switch thingType := thing.(type) {
	case Dog:
		dog := thing.(Dog)
		fmt.Printf("Good boy %s!\n", dog.Name)
	case Person:
		fmt.Printf("This thing is of type %T\n", thingType)
	default:
		fmt.Printf("We have an unknown thing...")
	}
}

func main() {

	thing1 := getIntThing()

	// Always check that the value we get is an int
	intVal1, ok := thing1.(int)
	if ok {
		intVal1++
		fmt.Printf("%d\n", intVal1)
	}

	thing2 := getFloatThing()
	_, ok = thing2.(int)
	if !ok {
		fmt.Println("Yikes! thing2 is a float thing!")
	}
	thing3 := getPersonThing()
	thing4 := getDogThing()

	doThingWithThing(thing3)
	doThingWithThing(thing4)

}
