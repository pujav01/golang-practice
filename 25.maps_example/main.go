package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 100) // create a slice
	for i, _ := range slice1 {
		slice1[i] = rand.Intn(50) // assigning values to the slice using rand
	}
	fmt.Println(slice1)

	mymap := make(map[int]uint16)

	for _, v := range slice1 {

		_, ok := mymap[v]
		if ok {
			mymap[v] = mymap[v] + 1
		} else {
			mymap[v] = 1
		}
	}
	fmt.Println(mymap)

}

/*

func Delete(mymap map[int]uint, key int)(bool, error) {
	if mymap == nil {
		return false, errors.New("map is nill")
	}
	if _, ok := mymap[key]; ok {
		Delete(mymap, key)

	}

}
*/
