package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	txt := "Hello World I am learning GO"
	fmt.Println("Count of vowels:", countVowels(txt))

}
