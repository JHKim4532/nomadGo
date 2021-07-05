package DataStructure

import "fmt"

func arrAndslice() {
	namesarr := [5]string{"nico", "lynn", "dal"}
	namesarr[3] = "alala"
	namesslice := []string{"nico", "lynn", "dal"}
	namesslice = append(namesslice, "alala")
	fmt.Println(namesarr, namesslice)
}
