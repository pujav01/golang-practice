package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	count := func(str string) (int, int, int) {
		vowels := "aeiouAEIOU"
		vowelCount := 0
		consonantCount := 0
		specialCharCount := 0

		for _, char := range str {
			if unicode.IsLetter(char) {
				if strings.ContainsRune(vowels, char) {
					vowelCount++
				} else {
					consonantCount++
				}
			} else if unicode.IsPunct(char) || unicode.IsSymbol(char) || unicode.IsSpace(char) {
				specialCharCount++
			}
		}
		return vowelCount, consonantCount, specialCharCount
	}

	input := "Hello World!"
	vowels, consonants, specialChars := count(input)
	fmt.Printf("vowels: %d, Consonants: %d, Special Characters: %d\n", vowels, consonants, specialChars)
}
