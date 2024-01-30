/*
*	26) Разработать программу, которая проверяет,
*	что все символы в строке уникальные (true — если уникальные, false etc).
*	Функция проверки должна быть регистронезависимой.
*
*	Например:
*	abcd — true
*	abCdefAaf — false
*	aabcd — false
*
*************************************************/
package main

import (
	"fmt"
	"strings"
)

// проверяет, что все символы в строке уникальные
func hasUnique(s string) bool {
	hash := make(map[rune]struct{})
	ls := strings.ToLower(s)

	for _, c := range ls {
		if _, ok := hash[c]; ok {
			return false
		}
		hash[c] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(hasUnique("abcd"))
	fmt.Println(hasUnique("abCdefAaf"))
	fmt.Println(hasUnique("aabcd"))
}
