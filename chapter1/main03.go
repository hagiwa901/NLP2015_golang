/*
"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
*/
package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	var s1 string
	// ,の削除
	str = strings.Replace(str, ",", "", -1)
	// .の削除
	str = strings.Replace(str, ".", "", -1)
	// 空白で区切る
	slice := strings.Split(str, " ")
	for _, s := range slice {
		s1 += s[0:1]
	}
	fmt.Println(s1)
}
