/*
"Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，
それ以外の単語は先頭に2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
*/

package main

import (
	"strings"
	"fmt"
)

/*
https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
*/
func contains(i int, num []int) bool {
	for _, n := range num {
		if n == i {
			return true
		}
	}
	return false
}

func main() {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	// ,の削除
	str = strings.Replace(str, ",", "", -1)
	// .の削除
	str = strings.Replace(str, ".", "", -1)
	// 空白で区切る
	slice := strings.Split(str, " ")
	// 値の格納
	m := map[int]string{}
	// 番号
	num := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}

	for i, s := range slice {
		if contains(i + 1, num) {
			m[i + 1] = s[0:1]
		} else {
			m[i + 1] = s[0:2]
		}
		
	}
	fmt.Println(m)
}
