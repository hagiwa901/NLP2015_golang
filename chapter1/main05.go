package main

import (
	"strings"
	"fmt"
)

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

	for i, s := range slice {
		if i % 2 == 0 {
			m[i + 1] = s[0:1]
		} else {
			m[i + 1] = s[1:2]
		}
		
	}
	fmt.Println(m)
}
