/*
与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
・英小文字ならば(219 - 文字コード)の文字に置換
・その他の文字はそのまま出力
この関数を用い，英語のメッセージを暗号化・復号化せよ．
*/
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
