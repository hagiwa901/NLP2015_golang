package main

import (
	"fmt"
)

func main() {
	/*
	https://itsakura.com/golang-slice
	*/
	str1 := "パトカー"
	str2 := "タクシー"
	var str string
	const char_num = 3
	// UTF8で書かれているため、1文字3バイトずつ動かす
	for i := 0; i < len(str1); i += char_num{
		str += str1[i:i+char_num] + str2[i:i+char_num]
	}
	fmt.Println(str)
}
