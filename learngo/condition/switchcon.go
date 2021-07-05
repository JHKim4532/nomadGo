package condition

import "fmt"

func canIDrinksw(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 19:
		return true
	}
	return false
}

func switchcon() {
	fmt.Println(canIDrink(20))
}
