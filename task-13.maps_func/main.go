package main

import (
	"errors"
	"fmt"
)

func main() {
	var mymap map[string]any
	mymap = make(map[string]any)
	mymap["522001"] = "Guntur-1"
	mymap["522002"] = "Guntur-2"
	mymap["560086"] = "Yeshvantpur Bangalore"
	mymap["560096"] = "Mahalakshmi Bangalore"
	mymap["690054"] = "Trivandrum-1"

	keyToDelete := "522001"
	boolean, err := Delete(mymap, keyToDelete)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		if boolean {
			fmt.Println("Key", keyToDelete, "deleted")
		} else {
			fmt.Println("Key", keyToDelete, "was not found")
		}
	}
	keyToUpdate := "522002"
	newValue := "Hyderabad"
	updated, err := Update(mymap, keyToUpdate, newValue)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		if updated {
			fmt.Println("Key", keyToUpdate, "updated")
		} else {
			fmt.Println("Key", keyToUpdate, "was not found")
		}

	}
	fmt.Println("updated map:", mymap)
}

func Delete(mymap map[string]any, key string) (bool, error) {
	if mymap == nil {
		return false, errors.New("map is nil")
	}

	_, exists := mymap[key]
	if exists {
		delete(mymap, key)
		return true, nil
	} else {
		return false, nil
	}
}

func Update(mymap map[string]any, key string, value any) (bool, error) {
	if mymap == nil {
		return false, errors.New("map is nil")
	}

	_, exists := mymap[key]
	if !exists {
		return false, errors.New("key does not exist")
	} else {
		mymap[key] = value
		return true, nil
	}
}
