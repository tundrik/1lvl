/*
*	19) Разработать программу, которая переворачивает подаваемую на вход строку
*	(например: «главрыба — абырвалг»). Символы могут быть unicode.
*
*************************************************/
package main

import (
	"fmt"
	"strings"
)

// переворачивает подаваемую на вход строку
func reverse(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	// идем в обратном порядке
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

func main() {
	fmt.Println(reverse("главрыба"))
	fmt.Println(reverse("狐犬"))
	fmt.Println(reverse("😎⚽"))
}
