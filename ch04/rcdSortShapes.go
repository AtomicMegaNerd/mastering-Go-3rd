package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	min = 1
	max = 5
)

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

type Sphere struct {
	r float64
}

func (s Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * s.r * s.r * s.r
}

type shapes []Shape3D

func (s shapes) Len() int {
	return len(s)
}

func (s shapes) Less(i, j int) bool {
	return s[i].Vol() < s[j].Vol()
}

func (s shapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func PrintShapes(s shapes) {
	for _, v := range s {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}

func main() {

	data := shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}

		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}
	PrintShapes(data)

	sort.Sort(data)
	PrintShapes(data)

	sort.Sort(sort.Reverse(data))
	PrintShapes(data)
}
