package basicfunction

import (
	"fmt"
	"strings"
)

func lenAndUpperOthReturn(name string) (length int, uppercase string) {
	defer fmt.Println("lAU func done") // defer
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
} //naked return

func func2() {
	totalLength, up := lenAndUpperOthReturn("nico")
	fmt.Println(totalLength, up)
}
