package main

import (
	"fmt"

	load_dictionary "palindromes/load"
)

func main() {
	word_list := load_dictionary.Load("2of4brif.txt")

	var pali_list []string

	for _, word := range word_list {
		if len(word) > 1 && word == reverse(word) {
			pali_list = append(pali_list, word)
		}
	}

	fmt.Printf("\nNumber of palindromes found = %d\n", len(pali_list))
	fmt.Print(pali_list)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
