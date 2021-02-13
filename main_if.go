package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// naked return
func lenAndUpper2(name string) (length int, uppdercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppdercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words[1])
	fmt.Println(words)

}

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	// for idx, number := range numbers {
	// 	fmt.Println(idx, number)
	// }
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
func main() {
	// const name string= "dani"

	// var name string = "dani"
	// name := "dani"
	// totalLength, uppderName := lenAndUpper(name)
	// fmt.Println(totalLength, uppderName)
	// fmt.Println(multiply(2, 2))

	// repeatMe("nico", "dani", "haesoo", "marl", "flynn")
	// lenAndUpper2(name)
	// total := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(total)
	fmt.Println(canIDrink(18))
}
