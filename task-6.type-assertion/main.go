package main

import "fmt"

func main() {
	// integer types
	var intVar int = 1
	var int8Var int8 = 2
	var int16Var int16 = 3
	var int32Var int32 = 4
	var int64Var int64 = 5

	// unsigned integer types
	var uintVar uint = 10
	var uint8Var uint8 = 11
	var uint16Var uint16 = 12
	var uint32Var uint32 = 13
	var uint64Var uint64 = 14

	// floating point types
	var float32Var float32 = 15.8
	var float64Var float64 = 20.7

	// rune type (alias for int32)
	var runeVar1 rune = 'A'

	// byte type (alias for uint8)
	var byteVar byte = 255

	var iface1 any
	iface1 = 5

	var iface2 any
	iface2 = 5.8

	addResult := float64(intVar) + float64(int8Var) + float64(int16Var) + float64(int32Var) + float64(int64Var) + float64(uintVar) + float64(uint8Var) + float64(uint16Var) + float64(uint32Var) + float64(uint64Var) + float64(float32Var) + float64(float64Var) + float64(runeVar1) + float64(byteVar) + float64(iface1.(int)) + iface2.(float64)
	fmt.Println("Addition:", addResult)

	mulResult := float64(intVar) * float64(int8Var) * float64(int16Var) * float64(int32Var) * float64(int64Var) * float64(uintVar) * float64(uint8Var) * float64(uint16Var) * float64(uint32Var) * float64(uint64Var) * float64(float32Var) * float64(float64Var) * float64(runeVar1) * float64(byteVar) * float64(iface1.(int)) * iface2.(float64)
	fmt.Println("Multiplication:", mulResult)

}
