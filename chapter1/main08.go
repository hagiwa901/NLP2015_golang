package main

import (
	"fmt"
)

func cipher(str string) string {
	var s string
	runes := []rune(str)
	for _, r := range runes {
		if(r >= 'a' && r <= 'z'){
			s += string(219 - r)
		} else {
			s += string(r)
		}
	}
	return s
}

func main() {
	str := "abc"
	fmt.Println(cipher(str))
}
