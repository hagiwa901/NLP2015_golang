package main

import (
	"strings"
	"fmt"
)

func ngram(str string, n int) {
	// 単語
	slice := strings.Split(str, " ")
	size := len(slice)
	for i := 0; i < size; i += n {
		if i + n < size {
			fmt.Println(slice[i:i+n])
		} else {
			fmt.Println(slice[i:])
		}
	}

	// 文字
	size = len(str)
	for i := 0; i < size; i += n {
		if i + n < size {
			fmt.Println(str[i:i+n])
		} else {
			fmt.Println(str[i:])
		}
	}
}

func main() {
	str := "I am an NLPer"

	fmt.Println(str)
	ngram(str, 2)
}
