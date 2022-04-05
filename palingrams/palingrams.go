package main

import (
	"fmt"

	load_dictionary "palindromes/load"
)

var word_list = load_dictionary.Load("2of4brif.txt")

func main() {
	palingrams_sorted := find_palingrams()
	// sort.Strings(sorted_palingrams)
	fmt.Printf("\nNumber of palingrams = %d\n", len(palingrams_sorted))
	for _, v := range palingrams_sorted {
		fmt.Println(v)
	}
}

func find_palingrams() []string {
	var pali_list []string
	for _, word := range word_list {
		end := len(word)
		rev_word := reverse(word)
		if end > 1 {
			for i := 0; i < end; i++ {
				if word[i:] == rev_word[:end-i] && in(rev_word[end-i:], word_list) {
					pali_list = append(pali_list, word+" "+rev_word[end-i:])
				}
				if word[:i] == rev_word[end-i:] && in(rev_word[:end-i], word_list) {
					pali_list = append(pali_list, rev_word[:end-i]+" "+word)
				}
			}
		}
	}
	return pali_list
}

func in(value string, array []string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
