package main

import (
	"fmt"

	"github.com/sssunda/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	// Search
	// fmt.Println(dictionary["first"])
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	baseword := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search("hello")
	// fmt.Println(hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)

	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
