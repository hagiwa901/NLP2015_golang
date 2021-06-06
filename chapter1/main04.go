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
