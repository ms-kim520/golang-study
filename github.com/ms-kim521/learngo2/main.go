package main

import (
	"fmt"

	"github.com/github.com/ms-kim521/learngo2/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
