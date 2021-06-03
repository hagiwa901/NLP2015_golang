package main

import (
	"fmt"
)

func main() {
	/*
	https://itsakura.com/golang-slice
	*/
	str := "パタトクカシーー"
	const char_num = 3
	// UTF8で書かれているため、1文字3バイトずつ動かす
	for i := 0; i < len(str); i += char_num * 2{
		fmt.Print(str[i:i+3])
	}
	fmt.Println()
}
