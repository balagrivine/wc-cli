package main

import (
	"fmt"
)

func main() {
	toEncode := "aaabbbbrrr"

	charCount := make(map[string]int, 10)
	var encoded string

	// var encoded string
	for i := range toEncode {
		// Current element accessed as byte, convert to string
		curr := string(toEncode[i])
		_, ok := charCount[curr]
		if ok {
			charCount[curr] += 1
		} else {
			charCount[curr] = 1
		}
	}
	fmt.Println(toEncode)
	for k, v := range charCount {
		encoded += fmt.Sprintf("%s%d", k, v)
	}
	fmt.Println(encoded)
	fmt.Println(charCount)
}
