package main

import (
	"fmt"

	"github.com/JHKim4532/nomadGo/GoDict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	sdictionary := mydict.Dictionary{}
	word := "hello"
	sdef := "Greeting"
	adderr := sdictionary.Add(word, sdef)
	if adderr != nil {
		fmt.Println(adderr)
	}
	hello, _ := sdictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	adderr = sdictionary.Add(word, sdef)
	if adderr != nil {
		fmt.Println(adderr)
	}

	tdictionary := mydict.Dictionary{}
	baseword := "hello"
	tdictionary.Add(baseword, "First")
	upderr := tdictionary.Update("nono", "Second")
	if upderr != nil {
		fmt.Println(upderr)
	}
	word, _ = tdictionary.Search(baseword)
	fmt.Println(word)

	delerr := tdictionary.Delete("nonono")
	if delerr != nil {
		fmt.Println(delerr)
	}
	word, _ = tdictionary.Search(baseword)
	fmt.Println(word)
}
