package main

import "fmt"

func main() {

	company := Company{
		Id:        1,
		Name:      "company1",
		Website:   "www.website.com",
		Telephone: "1234567890",
		Fax:       "1234",
		Address: Address{
			City:    "city1",
			Line1:   "line1",
			Street:  "street1",
			State:   "state1",
			Country: "country1",
			Pincode: "5678",
			Map: Map{
				Lan: 12.0972,
				Lat: -86.782,
			},
		},
	}
	lan := company.Address.Map.Lan
	lat := company.Address.Map.Lat

	fmt.Printf("Lan: %f, Lat: %f\n", lan, lat)
}

type Map struct {
	Lan float32
	Lat float32
}

type Company struct {
	Id        int
	Name      string
	Website   string
	Telephone string
	Fax       string
	Address
}

type Address struct {
	City    string
	Line1   string
	Street  string
	State   string
	Country string
	Pincode string
	Map
}
