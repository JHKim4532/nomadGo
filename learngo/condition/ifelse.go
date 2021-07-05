package condition

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge <= 18 {
		return false
	}
	return true
}

func ifelse() {
	fmt.Println(canIDrink(17))
}
