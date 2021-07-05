package basicfunction

import (
	"fmt"
	"strings"
)

func repeatMe(name ...string) {
	fmt.Println(name)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func funcExam() {
	fmt.Println(multiply(2, 2))

	totalLength, nameUpper := lenAndUpper("Name")
	total2, _ := lenAndUpper("Test")
	fmt.Println(total2)
	fmt.Println(totalLength, nameUpper)

	repeatMe("hong", "hun", "yeong")
}
