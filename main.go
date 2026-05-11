package main

import (
	"fmt"
	"strings"
)

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func IsPalindrome(str string) bool {
	runes := []rune(str)
	var builder strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	return str == builder.String()
}

func IsAutomorphic(num int) bool {
	sq := num * num

	for num > 0 {
		if sq%10 != num%10 {
			return false
		}
		num /= 10
		sq /= 10
	}
	return true
}

func ReverseNum(num int) int {
	sum := 0
	for num > 0 {
		sum = sum*10 + (num % 10)
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(LCM(12, 18))
	fmt.Println(IsPalindrome("malayalam"))
	fmt.Println(IsAutomorphic(25))
	fmt.Println(ReverseNum(19481))
}
