package variables

import "fmt"

func pack() {
	name := "hong" // cannot use this out of func
	var sname string = "kjh"
	const tname string = "test"

	fmt.Println(name)
	fmt.Println(sname)
}
