package main

import "fmt"

func main() {
	var mymap map[string]any
	mymap = make(map[string]any)
	mymap["522001"] = "Guntur-1"
	mymap["522002"] = "Guntur-2"
	mymap["560086"] = "Bangalore"
	mymap["560096"] = "Mysore"

	for key, value := range mymap {
		fmt.Println("key:", key, "value:", value)
	}

	_, ok := mymap["522001"]
	if ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exists")
	}

	delete(mymap, "522001") // deletes the key from the map

	for key, value := range mymap {
		fmt.Println("key:", key, "value:", value)
	}

}
