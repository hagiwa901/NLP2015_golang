/*
「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．
*/
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
		fmt.Print(str[i:i+char_num])
	}
	fmt.Println()
}
