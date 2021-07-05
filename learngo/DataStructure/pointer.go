package DataStructure

import "fmt"

func pointer() {
	a := 2
	b := &a
	*b = 2020
	fmt.Println(a)
}
