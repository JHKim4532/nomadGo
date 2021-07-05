package condition

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { //first argu is index
		total += number
	}
	/*
		for i := 0; i < len(numbers); i++{
			fmt.Println(numbers[i])
		}
	*/
	return total
}

func looping() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
