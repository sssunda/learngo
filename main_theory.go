package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// a := 2
	// b := &a
	// a = 10
	// fmt.Println(a, *b)
	// names := [5]string{"nico", "lynn", "dal"}
	// names := []string{"nico", "lynn", "dal"}
	// names = append(names, "3")
	// nico := map[string]string{"name": "nico", "age": "12"}
	favFood := []string{"kimchi", "bulgogi"}
	nico := person{name: "nico", age: 12, favFood: favFood}

	fmt.Println(nico)

}
