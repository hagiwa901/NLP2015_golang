package main

import (
	"fmt"
)

/*
https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
How to reverse a string in Go?
*/
// string型を逆順にする
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "stressed"
	fmt.Println(Reverse(str))
}
