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

	// bool type
	var boolVar1 bool = true
	var boolVar2 bool = false

	// string type
	var stringVar1 string = "Go "
	var stringVar2 string = "Programming"

	// rune type (alias for int32)
	var runeVar1 rune = 'A'
	var runeVar2 rune = 'B'

	// byte type (alias for uint8)
	var byteVar byte = 255

	intAdd := intVar + int(int8Var) + int(int16Var) + int(int32Var) + int(int64Var)
	fmt.Println("int addition result:", intAdd)

	intMul := intVar * int(int8Var) * int(int16Var) * int(int32Var) * int(int64Var)
	fmt.Println("int multiplication result:", intMul)

	uintAdd := uintVar + uint(uint8Var) + uint(uint16Var) + uint(uint32Var) + uint(uint64Var)
	fmt.Println("uint addition result:", uintAdd)

	uintMul := uintVar * uint(uint8Var) * uint(uint16Var) * uint(uint32Var) * uint(uint64Var)
	fmt.Println("uint multiplication result:", uintMul)

	floatAdd := float32Var + float32(float64Var)
	fmt.Println("float addition result:", floatAdd)

	floatMul := float32Var * float32(float64Var)
	fmt.Println("float multiplication result:", floatMul)

	boolAnd := boolVar1 && boolVar2
	fmt.Println("And operation on boolean:", boolAnd)

	boolOr := boolVar1 || boolVar2
	fmt.Println("Or operation on boolean:", boolOr)

	stringAdd := stringVar1 + stringVar2
	fmt.Println("Addition operation on string:", stringAdd)

	runeAdd := runeVar1 + runeVar2
	fmt.Println("Addition operation on rune:", runeAdd)

	runeMul := runeVar1 * runeVar2
	fmt.Println("Multiplication operation on rune:", runeMul)

	fmt.Println("byte:", byteVar)

}
